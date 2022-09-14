package impl

import (
	"context"
	"fmt"
	"restful-api-demo/apps/host"
)

// 把Host保存数据库
func (i *HostServiceImpl) save(ctx context.Context, ins *host.Host) error {

	var err error
	tx, err := i.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("start tx error, %s", err)
	}
	defer func() {
		if err != nil {
			if err := tx.Rollback(); err != nil {
				i.l.Errorf("rollback error, %s", err)
			}
		} else {
			if err := tx.Commit(); err != nil {
				i.l.Errorf("commit error, %s", err)
			}
		}
	}()
	// insert resource table
	rstmt, err := tx.Prepare(InsertResourceSQL)
	if err != nil {
		return err
	}
	defer rstmt.Close()
	_, err = rstmt.Exec(ins.Id, ins.Vendor, ins.Region, ins.CreateAt, ins.ExpireAt, ins.Type,
		ins.Name, ins.Description, ins.Status, ins.UpdateAt, ins.SyncAt, ins.Account, ins.PublicIP,
		ins.PrivateIP)
	if err != nil {
		return err
	}

	// insert describe table
	dstmt, err := tx.Prepare(InsertDescribeSQL)
	if err != nil {
		return err
	}
	defer dstmt.Close()
	_, err = dstmt.Exec(ins.Id, ins.CPU, ins.Memory, ins.GPUAmount, ins.GPUSpec,
		ins.OSType, ins.OSName, ins.SerialNumber)
	if err != nil {
		return err
	}
	return nil
}
