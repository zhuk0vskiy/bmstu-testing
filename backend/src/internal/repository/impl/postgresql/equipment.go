package postgresql

import (
	"backend/src/internal/model"
	"backend/src/internal/model/dto"
	repositoryInterface "backend/src/internal/repository/interface"
	"backend/src/pkg/time_parser"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type EquipmentRepository struct {
	db *pgxpool.Pool
	//db repositoryInterface.IPool
}

//func NewEquipmentRepository(db *pgxpool.Pool) repositoryInterface.IEquipmentRepository {
//	return &EquipmentRepository{
//		db: db,
//	}
//}

func NewEquipmentRepository(db *pgxpool.Pool) repositoryInterface.IEquipmentRepository {
	return &EquipmentRepository{
		db: db,
	}
}

func (r EquipmentRepository) GetByReserve(ctx context.Context, request *dto.GetEquipmentByReserveRequest) (equipments []*model.Equipment, err error) {
	query := `
				select equipment.id, 
				       equipment.name,
				       equipment.type,
				       equipment.studio_id 
				from equipment join reserved_equipments on equipment.id = reserved_equipments.equipment_id
				where reserved_equipments.reserve_id = $1`

	rows, err := r.db.Query(
		ctx,
		query,
		request.ReserveId,
	)

	if err != nil {
		return nil, fmt.Errorf("ошибка при запросе: %w", err)
	}

	equipments = make([]*model.Equipment, 0)

	for rows.Next() {
		tmp := new(model.Equipment)

		err = rows.Scan(
			&tmp.Id,
			&tmp.Name,
			&tmp.EquipmentType,
			&tmp.StudioId,
		)

		equipments = append(equipments, tmp)

		if err != nil {
			return nil, fmt.Errorf("сканирование полученных строк: %w", err)
		}
	}

	return equipments, err
}

func (r EquipmentRepository) Get(ctx context.Context, request *dto.GetEquipmentRequest) (equipment *model.Equipment, err error) {

	query := `    select id, 
				       name,
				       type,
				       studio_id
				from equipment
				where id = $1`

	equipment = new(model.Equipment)

	err = r.db.QueryRow(
		ctx,
		query,
		request.Id,
	).Scan(
		&equipment.Id,
		&equipment.Name,
		&equipment.EquipmentType,
		&equipment.StudioId,
	)

	if err != nil {
		return nil, fmt.Errorf("get запрос не выполнен: %w", err)
	}

	return equipment, err
}

func (r EquipmentRepository) GetByStudio(ctx context.Context, request *dto.GetEquipmentByStudioRequest) (equipments []*model.Equipment, err error) {
	query := `
				select equipment.id, 
				       equipment.name,
				       equipment.type,
				       equipment.studio_id 
				from equipment 
				where equipment.studio_id = $1`

	rows, err := r.db.Query(
		ctx,
		query,
		request.StudioId,
	)

	if err != nil {
		return nil, fmt.Errorf("ошибка при getAll запросе: %w", err)
	}

	equipments = make([]*model.Equipment, 0)

	for rows.Next() {
		tmp := new(model.Equipment)

		err = rows.Scan(
			&tmp.Id,
			&tmp.Name,
			&tmp.EquipmentType,
			&tmp.StudioId,
		)

		equipments = append(equipments, tmp)

		if err != nil {
			return nil, fmt.Errorf("сканирование полученных строк: %w", err)
		}
	}

	return equipments, err
}

func (r EquipmentRepository) GetFullTimeFreeByStudioAndType(ctx context.Context, request *dto.GetEquipmentFullTimeFreeByStudioAndTypeRequest) (equipments []*model.Equipment, err error) {
	//TODO: проверить
	query := `select equipment.id,
				       equipment.name,
				       equipment.type,
				       equipment.studio_id
				from equipment where equipment.studio_id = $1 and equipment.type = $2 and not exists
				    (select * from reserved_equipments where equipment.id = reserved_equipments.equipment_id)`
	//
	//query := `select equipment.id,
	//			       equipment.name,
	//			       equipment.type,
	//			       equipment.studio_id
	//			from equipment`
	rows, err := r.db.Query(
		ctx,
		query,
		request.StudioId,
		request.Type,
	)

	if err != nil {
		return nil, fmt.Errorf("ошибка при запросе: %w", err)
	}

	equipments = make([]*model.Equipment, 0)

	for rows.Next() {
		tmp := new(model.Equipment)

		err = rows.Scan(
			&tmp.Id,
			&tmp.Name,
			&tmp.EquipmentType,
			&tmp.StudioId,
		)

		equipments = append(equipments, tmp)

		if err != nil {
			return nil, fmt.Errorf("сканирование полученных строк: %w", err)
		}
	}

	return equipments, err
}

func (r EquipmentRepository) GetNotFullTimeFreeByStudioAndType(ctx context.Context, request *dto.GetEquipmentNotFullTimeFreeByStudioAndTypeRequest) (equipmentsAndTime []*dto.EquipmentAndTime, err error) {
	//TODO: доделать
	startTimeh := time.Now()
	query := `select equipment.id, 
				       equipment.name,
				       equipment.type,
				       equipment.studio_id,
					   to_char(reserve.start_time, 'YYYY-MM-DD HH24:MI:SS'),
					   to_char(reserve.end_time, 'YYYY-MM-DD HH24:MI:SS')
				from equipment, reserve where equipment.studio_id = $1 and equipment.type = $2 and exists 
				    (select * from reserved_equipments where equipment.id = reserved_equipments.equipment_id)`
	rows, err := r.db.Query(
		ctx,
		query,
		request.StudioId,
		request.Type,
	)
	fmt.Println("time:", time.Since(startTimeh))
	if err != nil {
		return nil, fmt.Errorf("ошибка при запросе: %w", err)
	}

	equipmentsAndTime = make([]*dto.EquipmentAndTime, 0)
	var startTime string = ""
	var endTime string = ""

	for rows.Next() {
		var tmp dto.EquipmentAndTime
		tmp.TimeInterval = new(model.TimeInterval)
		tmp.Equipment = new(model.Equipment)
		err = rows.Scan(
			&tmp.Equipment.Id,
			&tmp.Equipment.Name,
			&tmp.Equipment.EquipmentType,
			&tmp.Equipment.StudioId,
			&startTime,
			&endTime,
		)

		if err != nil {
			return nil, fmt.Errorf("сканирование полученных строк: %w", err)
		}

		tmp.TimeInterval.StartTime, err = time_parser.StringToDate(startTime)
		if err != nil {
			return nil, fmt.Errorf("ошибка конвертирования: %w", err)
		}

		tmp.TimeInterval.EndTime, err = time_parser.StringToDate(endTime)
		if err != nil {
			return nil, fmt.Errorf("ошибка конвертирования: %w", err)
		}
		equipmentsAndTime = append(equipmentsAndTime, &tmp)

	}

	return equipmentsAndTime, err

}

func (r EquipmentRepository) Update(ctx context.Context, request *dto.UpdateEquipmentRequest) (err error) {
	query := `
			update equipment
			set 
			    id = $1,
			    name = $2,
				type = $3,
				studio_id = $4
			where id = $1`

	_, err = r.db.Exec(
		ctx,
		query,
		request.Id,
		request.Name,
		request.Type,
		request.StudioId,
	)
	if err != nil {
		return fmt.Errorf("обновление информации о компании: %w", err)
	}

	return err
}

func (r EquipmentRepository) Add(ctx context.Context, request *dto.AddEquipmentRequest) (err error) {
	query := `insert into equipment(name, type, studio_id) values ($1, $2, $3)`

	_, err = r.db.Exec(
		ctx,
		query,
		request.Name,
		request.Type,
		request.StudioId,
	)
	if err != nil {
		return fmt.Errorf("создание финансового отчета: %w", err)
	}

	return err
}

func (r EquipmentRepository) Delete(ctx context.Context, request *dto.DeleteEquipmentRequest) (err error) {
	query := `delete from equipment where id = $1`

	_, err = r.db.Exec(
		ctx,
		query,
		request.Id,
	)
	if err != nil {
		return fmt.Errorf("удаление отчета по id: %w", err)
	}

	return err
}
