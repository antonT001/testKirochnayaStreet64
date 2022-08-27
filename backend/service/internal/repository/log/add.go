package log

func (r *log) Add(stmt string, valueArgs []interface{}) error {

	err := r.db.Exec(stmt, valueArgs...)
	if err != nil {
		return err
	}

	return nil
}
