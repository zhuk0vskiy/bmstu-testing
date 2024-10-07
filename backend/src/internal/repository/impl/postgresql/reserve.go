package postgresql

import (
	"backend/src/internal/model"
	"backend/src/internal/model/dto"
	repositoryInterface "backend/src/internal/repository/interface"
	"backend/src/pkg/time_parser"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ReserveRepository struct {
	db *pgxpool.Pool
	//db repositoryInterface.IPool
}

func NewReserveRepository(db *pgxpool.Pool) repositoryInterface.IReserveRepository {
	return &ReserveRepository{
		db: db,
	}
}

func (r ReserveRepository) GetAll(ctx context.Context, request *dto.GetAllReserveRequest) (reserves []*model.Reserve, err error) {
	query := `
				select id, 
				       user_id,
				       room_id,
				       producer_id,
				       instrumentalist_id,
					   to_char(start_time, 'YYYY-MM-DD HH24:MI:SS'),
					   to_char(end_time, 'YYYY-MM-DD HH24:MI:SS')
				from reserve`

	rows, err := r.db.Query(
		ctx,
		query,
	)

	if err != nil {
		return nil, fmt.Errorf("ошибка при запросе: %w", err)
	}

	reserves = make([]*model.Reserve, 0)
	var startTime string = ""
	var endTime string = ""

	for rows.Next() {
		tmp := new(model.Reserve)
		tmp.TimeInterval = new(model.TimeInterval)
		err = rows.Scan(
			&tmp.Id,
			&tmp.UserId,
			&tmp.RoomId,
			&tmp.ProducerId,
			&tmp.InstrumentalistId,
			&startTime,
			&endTime,
		)
		if err != nil {
			return nil, fmt.Errorf("сканирование полученных строк: %w", err)
		}

		fmt.Println(startTime)

		tmp.TimeInterval.StartTime, err = time_parser.StringToDate(startTime)
		//fmt.Println(tmp.TimeInterval.StartTime)
		fmt.Println(4)
		if err != nil {
			return nil, fmt.Errorf("ошибка конвертирования: %w", err)
		}

		tmp.TimeInterval.EndTime, err = time_parser.StringToDate(endTime)
		if err != nil {
			return nil, fmt.Errorf("ошибка конвертирования: %w", err)
		}
		fmt.Println(1)
		reserves = append(reserves, tmp)

	}

	return reserves, err
}

func (r ReserveRepository) IsRoomReserve(ctx context.Context, request *dto.IsRoomReserveRequest) (isReserve bool, err error) {
	query := `
				select count(*)
				from reserve 
				where room_id = $1`

	var reserveCount int64

	err = r.db.QueryRow(
		ctx,
		query,
		request.RoomId,
	).Scan(
		&reserveCount,
	)

	if err != nil {
		return true, fmt.Errorf("запрос не выполнен: %w", err)
	}

	if reserveCount == 0 {
		return false, err
	}

	return true, err
}

func (r ReserveRepository) IsInstrumentalistReserve(ctx context.Context, request *dto.IsInstrumentalistReserveRequest) (isReserve bool, err error) {
	query := `
				select count(*)
				from reserve 
				where instrumentalist_id = $1`

	var reserveCount int64

	err = r.db.QueryRow(
		ctx,
		query,
		request.InstrumentalistId,
	).Scan(
		&reserveCount,
	)

	if err != nil {
		return true, fmt.Errorf("запрос не выполнен: %w", err)
	}

	if reserveCount == 0 {
		return false, err
	}

	return true, err
}

func (r ReserveRepository) IsProducerReserve(ctx context.Context, request *dto.IsProducerReserveRequest) (isReserve bool, err error) {
	query := `
				select count(*)
				from reserve 
				where producer_id = $1`

	var reserveCount int64

	err = r.db.QueryRow(
		ctx,
		query,
		request.ProducerId,
	).Scan(
		&reserveCount,
	)

	if err != nil {
		return true, fmt.Errorf("запрос не выполнен: %w", err)
	}

	if reserveCount == 0 {
		return false, err
	}

	return true, err
}

func (r ReserveRepository) IsEquipmentReserve(ctx context.Context, request *dto.IsEquipmentReserveRequest) (isReserve bool, err error) {
	query := `
				select count(*)
				from reserved_equipments 
				where equipment_id = $1`

	var reserveCount int64

	err = r.db.QueryRow(
		ctx,
		query,
		request.EquipmentId,
	).Scan(
		&reserveCount,
	)

	if err != nil {
		return true, fmt.Errorf("запрос не выполнен: %w", err)
	}

	if reserveCount == 0 {
		return false, err
	}

	return true, err
}

func (r ReserveRepository) GetByRoomId(ctx context.Context, request *dto.GetReserveByRoomIdRequest) (reserves []*model.Reserve, err error) {
	query := `
				select id, 
				       user_id,
				       room_id,
				       producer_id,
				       instrumentalist_id,
					   to_char(start_time, 'YYYY-MM-DD HH24:MI:SS'),
					   to_char(end_time, 'YYYY-MM-DD HH24:MI:SS')
				from reserve
				where room_id = $1`

	rows, err := r.db.Query(
		ctx,
		query,
		request.RoomId,
	)

	if err != nil {
		return nil, fmt.Errorf("ошибка при запросе: %w", err)
	}

	reserves = make([]*model.Reserve, 0)
	var startTime string = ""
	var endTime string = ""

	for rows.Next() {
		tmp := new(model.Reserve)
		tmp.TimeInterval = new(model.TimeInterval)
		err = rows.Scan(
			&tmp.Id,
			&tmp.UserId,
			&tmp.RoomId,
			&tmp.ProducerId,
			&tmp.InstrumentalistId,
			&startTime,
			&endTime,
		)
		if err != nil {
			return nil, fmt.Errorf("сканирование полученных строк: %w", err)
		}

		fmt.Println(startTime)

		tmp.TimeInterval.StartTime, err = time_parser.StringToDate(startTime)
		//fmt.Println(tmp.TimeInterval.StartTime)
		fmt.Println(4)
		if err != nil {
			return nil, fmt.Errorf("ошибка конвертирования: %w", err)
		}

		tmp.TimeInterval.EndTime, err = time_parser.StringToDate(endTime)
		if err != nil {
			return nil, fmt.Errorf("ошибка конвертирования: %w", err)
		}
		fmt.Println(1)
		reserves = append(reserves, tmp)

	}

	return reserves, err
}

func (r ReserveRepository) GetByInstrumentalistId(ctx context.Context, request *dto.GetReserveByInstrumentalistIdRequest) (reserves []*model.Reserve, err error) {
	query := `
				select id, 
				       user_id,
				       room_id,
				       producer_id,
				       instrumentalist_id,
					   to_char(start_time, 'YYYY-MM-DD HH24:MI:SS'),
					   to_char(end_time, 'YYYY-MM-DD HH24:MI:SS')
				from reserve
				where instrumentalist_id = $1`

	rows, err := r.db.Query(
		ctx,
		query,
		request.InstrumentalistId,
	)

	if err != nil {
		return nil, fmt.Errorf("ошибка при запросе: %w", err)
	}

	reserves = make([]*model.Reserve, 0)
	var startTime string = ""
	var endTime string = ""

	for rows.Next() {
		tmp := new(model.Reserve)
		tmp.TimeInterval = new(model.TimeInterval)
		err = rows.Scan(
			&tmp.Id,
			&tmp.UserId,
			&tmp.RoomId,
			&tmp.ProducerId,
			&tmp.InstrumentalistId,
			&startTime,
			&endTime,
		)
		if err != nil {
			return nil, fmt.Errorf("сканирование полученных строк: %w", err)
		}

		fmt.Println(startTime)

		tmp.TimeInterval.StartTime, err = time_parser.StringToDate(startTime)
		//fmt.Println(tmp.TimeInterval.StartTime)
		fmt.Println(4)
		if err != nil {
			return nil, fmt.Errorf("ошибка конвертирования: %w", err)
		}

		tmp.TimeInterval.EndTime, err = time_parser.StringToDate(endTime)
		if err != nil {
			return nil, fmt.Errorf("ошибка конвертирования: %w", err)
		}
		fmt.Println(1)
		reserves = append(reserves, tmp)

	}

	return reserves, err
}

func (r ReserveRepository) GetByProducerId(ctx context.Context, request *dto.GetReserveByProducerIdRequest) (reserves []*model.Reserve, err error) {
	query := `
				select id, 
				       user_id,
				       room_id,
				       producer_id,
				       instrumentalist_id,
					   to_char(start_time, 'YYYY-MM-DD HH24:MI:SS'),
					   to_char(end_time, 'YYYY-MM-DD HH24:MI:SS')
				from reserve
				where producer_id = $1`

	rows, err := r.db.Query(
		ctx,
		query,
		request.ProducerId,
	)

	if err != nil {
		return nil, fmt.Errorf("ошибка при запросе: %w", err)
	}

	reserves = make([]*model.Reserve, 0)
	var startTime string = ""
	var endTime string = ""

	for rows.Next() {
		tmp := new(model.Reserve)
		tmp.TimeInterval = new(model.TimeInterval)
		err = rows.Scan(
			&tmp.Id,
			&tmp.UserId,
			&tmp.RoomId,
			&tmp.ProducerId,
			&tmp.InstrumentalistId,
			&startTime,
			&endTime,
		)
		if err != nil {
			return nil, fmt.Errorf("сканирование полученных строк: %w", err)
		}

		fmt.Println(startTime)

		tmp.TimeInterval.StartTime, err = time_parser.StringToDate(startTime)
		//fmt.Println(tmp.TimeInterval.StartTime)
		fmt.Println(4)
		if err != nil {
			return nil, fmt.Errorf("ошибка конвертирования: %w", err)
		}

		tmp.TimeInterval.EndTime, err = time_parser.StringToDate(endTime)
		if err != nil {
			return nil, fmt.Errorf("ошибка конвертирования: %w", err)
		}
		fmt.Println(1)
		reserves = append(reserves, tmp)

	}

	return reserves, err
}

func (r *ReserveRepository) GetUserReserves(ctx context.Context, request *dto.GetUserReservesRequest) (reserves []*model.Reserve, err error) {
	query := `
				select id, 
				       user_id,
				       room_id,
				       producer_id,
				       instrumentalist_id,
					   to_char(start_time, 'YYYY-MM-DD HH24:MI:SS'),
					   to_char(end_time, 'YYYY-MM-DD HH24:MI:SS')
				from reserve
				where user_id = $1`

	rows, err := r.db.Query(
		ctx,
		query,
		request.Id,
	)

	if err != nil {
		return nil, fmt.Errorf("ошибка при запросе: %w", err)
	}

	reserves = make([]*model.Reserve, 0)
	var startTime string = ""
	var endTime string = ""

	for rows.Next() {
		tmp := new(model.Reserve)
		tmp.TimeInterval = new(model.TimeInterval)
		err = rows.Scan(
			&tmp.Id,
			&tmp.UserId,
			&tmp.RoomId,
			&tmp.ProducerId,
			&tmp.InstrumentalistId,
			&startTime,
			&endTime,
		)
		if err != nil {
			return nil, fmt.Errorf("сканирование полученных строк: %w", err)
		}

		//fmt.Println(startTime)

		tmp.TimeInterval.StartTime, err = time_parser.StringToDate(startTime)
		//fmt.Println(tmp.TimeInterval.StartTime)
		//fmt.Println(4)
		if err != nil {
			return nil, fmt.Errorf("ошибка конвертирования: %w", err)
		}

		tmp.TimeInterval.EndTime, err = time_parser.StringToDate(endTime)
		if err != nil {
			return nil, fmt.Errorf("ошибка конвертирования: %w", err)
		}
		//fmt.Println(1)
		reserves = append(reserves, tmp)

	}

	return reserves, err
}

func (r *ReserveRepository) Add(ctx context.Context, request *dto.AddReserveRequest) (err error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("начало транзакции: %w", err)
	}
	defer func() {
		if err != nil {
			rollbackErr := tx.Rollback(ctx)
			if rollbackErr != nil {
				err = fmt.Errorf("rollback error: %w", rollbackErr)
			}
		}
	}()

	var reserveCount int64

	query := `select is_reserve($1, $2, $3, $4, $5, $6)`

	err = tx.QueryRow(
		ctx,
		query,
		request.UserId,
		request.RoomId,
		request.ProducerId,
		request.InstrumentalistId,
		request.TimeInterval.StartTime.Format("2006-01-02 15:04:05"),
		request.TimeInterval.EndTime.Format("2006-01-02 15:04:05"),
	).Scan(
		&reserveCount,
	)

	if err != nil {
		return fmt.Errorf("запрос не выполнен: %w", err)
	}

	if reserveCount != 0 {
		return fmt.Errorf("занято")
	}

	query = `insert into reserve(user_id, room_id, producer_id, instrumentalist_id, start_time, end_time) values ($1, $2, $3, $4, $5, $6) returning id`

	var reserveId int64

	err = tx.QueryRow(
		ctx,
		query,
		request.UserId,
		request.RoomId,
		request.ProducerId,
		request.InstrumentalistId,
		request.TimeInterval.StartTime.Format("2006-01-02 15:04:05"),
		request.TimeInterval.EndTime.Format("2006-01-02 15:04:05"),
	).Scan(&reserveId)
	if err != nil {
		return fmt.Errorf("создание брони: %w", err)
	}

	for _, equipmentId := range request.EquipmentId {
		query = `insert into reserved_equipments(reserve_id, equipment_id) values ($1, $2)`
		_, err = tx.Exec(
			ctx,
			query,
			reserveId,
			equipmentId,
		)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("commit tx: %w", err)
	}

	return err
}

func (r ReserveRepository) Delete(ctx context.Context, request *dto.DeleteReserveRequest) (err error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("начало транзакции: %w", err)
	}
	defer func() {
		if err != nil {
			rollbackErr := tx.Rollback(ctx)
			if rollbackErr != nil {
				err = fmt.Errorf("rollback error: %w", rollbackErr)
			}
		}
	}()

	query := `delete from reserved_equipments where reserve_id = $1`

	_, err = tx.Exec(
		ctx,
		query,
		request.Id,
	)
	if err != nil {
		return fmt.Errorf("удаление брони по id из reserved_equipments: %w", err)
	}

	query = `delete from reserve where id = $1`

	_, err = tx.Exec(
		ctx,
		query,
		request.Id,
	)
	if err != nil {
		return fmt.Errorf("удаление брони по id: %w", err)
	}
	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("commit tx: %w", err)
	}

	return err
}
