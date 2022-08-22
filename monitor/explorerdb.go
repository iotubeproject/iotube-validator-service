package monitor

import (
	"database/sql"
	"fmt"

	"github.com/crosschained-io/crosschained-service/types"
	"github.com/pkg/errors"
)

type ExplorerDB struct {
	db        *sql.DB
	tableName string
}

func NewExplorerDB(url, tableName string) (*ExplorerDB, error) {
	db, err := sql.Open("mysql", url)
	if err != nil {
		return nil, err
	}

	return &ExplorerDB{
		db:        db,
		tableName: tableName,
	}, nil
}

func (edb *ExplorerDB) Init() error {
	_, err := edb.db.Exec(fmt.Sprintf(
		"CREATE TABLE IF NOT EXISTS %s ("+
			"`cashier` varchar(42) NOT NULL,"+
			"`token` varchar(42) NOT NULL,"+
			"`tidx` bigint(20) NOT NULL,"+
			"`sender` varchar(42) NOT NULL,"+
			"`recipient` varchar(42) NOT NULL,"+
			"`amount` varchar(78) NOT NULL,"+
			"`fee` varchar(78),"+
			"`id` varchar(66) NOT NULL,"+
			"`creationTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,"+
			"`updateTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,"+
			"`status` varchar(10) NOT NULL DEFAULT 'new',"+
			"`txHash` varchar(66) DEFAULT NULL,"+
			"`txTimestamp` timestamp DEFAULT CURRENT_TIMESTAMP,"+
			"`gas` bigint(20),"+
			"`nonce` bigint(20),"+
			"`relayer` varchar(42) DEFAULT NULL,"+
			"`gasPrice` varchar(78) DEFAULT NULL,"+
			"`notes` varchar(45) DEFAULT NULL,"+
			"PRIMARY KEY (`cashier`,`token`,`tidx`),"+
			"UNIQUE KEY `id_UNIQUE` (`id`),"+
			"KEY `cashier_index` (`cashier`),"+
			"KEY `token_index` (`token`),"+
			"KEY `sender_index` (`sender`),"+
			"KEY `recipient_index` (`recipient`),"+
			"KEY `status_index` (`status`),"+
			"KEY `txHash_index` (`txHash`)"+
			") ENGINE=InnoDB DEFAULT CHARSET=latin1;",
		edb.tableName,
	))
	return err
}

func (edb *ExplorerDB) Backup(tasks []types.RelayTask, statuses, txHashes []string) error {
	if len(tasks) == 0 {
		return nil
	}
	if len(tasks) != len(statuses) || len(tasks) != len(txHashes) {
		return errors.New("invalid parameters")
	}
	tx, err := edb.db.Begin()
	if err != nil {
		return err
	}
	command := fmt.Sprintf(
		"INSERT IGNORE INTO %s (cashier, token, tidx, sender, recipient, amount, fee, id, status, txHash) "+
			"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?) "+
			"ON DUPLICATE KEY UPDATE status = VALUES(status), txHash = VALUES(txHash)",
		edb.tableName,
	)
	for i, task := range tasks {
		if _, err := tx.Exec(
			command,
			fmt.Sprintf("0x%040x", task.Deposit.DestinationTubeID),
			task.Deposit.DestinationToken,
			task.Deposit.SourceTubeNonce,
			task.Deposit.Sender,
			task.Deposit.Recipient,
			task.Deposit.Amount.String(),
			task.Fee.BigInt().String(),
			task.ID,
			statuses[i],
			txHashes[i],
		); err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				return rollbackErr
			}
			return err
		}
	}
	return tx.Commit()
}
