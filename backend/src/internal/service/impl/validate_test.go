package impl

//func TestValidateTimeService_GetSuitableStuff(t *testing.T) {
//	type fields struct {
//		roomRepo              repoInterface.IRoomRepository
//		equipmentRepo         repoInterface.IEquipmentRepository
//		producerRepo          repoInterface.IProducerRepository
//		instrumentalistRepo   repoInterface.IInstrumentalistRepository
//		reserveRepo           repoInterface.IReserveRepository
//		reservedEquipmentRepo repoInterface.IReservedEquipmentRepository
//	}
//	type args struct {
//		request *dto.GetSuitableStuffRequest
//	}
//	tests := []struct {
//		name                            string
//		fields                          fields
//		args                            args
//		wantNotReservedRooms            []*model.Room
//		wantNotReservedEquipments       [][]*model.Equipment
//		wantNotReservedProducers        []*model.Producer
//		wantNotReservedInstrumentalists []*model.Instrumentalist
//		wantErr                         bool
//	}{
//		//{
//		//	name: "test_pos_01",
//		//	args: args{
//		//
//		//		request: &dto.GetSuitableStuffRequest{
//		//			ChoosenInterval: &model.TimeInterval {
//		//StartTime: time.Date(2024, 4, 8, 13, 00, 00, 00, time.UTC),
//		//			EndTime:   time.Date(2024, 4, 8, 15, 00, 00, 00, time.UTC),
//		//			StudioId:  1,
//		//		},
//		//	},
//		//	wantErr: false,
//		//	wantNotReservedRooms: []*model.Room{
//		//		&model.Room{
//		//			Id:        1,
//		//			Name:      "1",
//		//			StudioId:  1,
//		//			StartHour: 9,
//		//			EndHour:   21,
//		//		},
//		//	},
//		//	wantNotReservedEquipments: [][]*model.Equipment{
//		//		{
//		//			&model.Equipment{
//		//				Id:            1,
//		//				Name:          "1",
//		//				StudioId:      1,
//		//				EquipmentType: 1,
//		//			},
//		//		},
//		//	},
//		//
//		//	wantNotReservedProducers: []*model.Producer{
//		//		&model.Producer{
//		//			Id:        1,
//		//			Name:      "1",
//		//			StudioId:  1,
//		//			StartHour: 9,
//		//			EndHour:   21,
//		//		},
//		//	},
//		//	wantNotReservedInstrumentalists: []*model.Instrumentalist{
//		//		&model.Instrumentalist{
//		//			Id:        1,
//		//			Name:      "1",
//		//			StudioId:  1,
//		//			StartHour: 9,
//		//			EndHour:   21,
//		//		},
//		//	},
//		//},
//	}
//	//roomRepo := new(pool.IRoomRepository)
//	//equipmentRepo := new(pool.IEquipmentRepository)
//	//producerRepo := new(pool.IProducerRepository)
//	//instrumentalistRepo := new(pool.IInstrumentalistRepository)
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			s := ValidateTimeService{
//				roomRepo:              tt.fields.roomRepo,
//				equipmentRepo:         tt.fields.equipmentRepo,
//				producerRepo:          tt.fields.producerRepo,
//				instrumentalistRepo:   tt.fields.instrumentalistRepo,
//				reserveRepo:           tt.fields.reserveRepo,
//				reservedEquipmentRepo: tt.fields.reservedEquipmentRepo,
//			}
//			gotNotReservedRooms, gotNotReservedEquipments, gotNotReservedProducers, gotNotReservedInstrumentalists, err := s.GetSuitableStuff(tt.args.request)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("GetSuitableStuff() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(gotNotReservedRooms, tt.wantNotReservedRooms) {
//				t.Errorf("GetSuitableStuff() gotNotReservedRooms = %v, want %v", gotNotReservedRooms, tt.wantNotReservedRooms)
//			}
//			if !reflect.DeepEqual(gotNotReservedEquipments, tt.wantNotReservedEquipments) {
//				t.Errorf("GetSuitableStuff() gotNotReservedEquipments = %v, want %v", gotNotReservedEquipments, tt.wantNotReservedEquipments)
//			}
//			if !reflect.DeepEqual(gotNotReservedProducers, tt.wantNotReservedProducers) {
//				t.Errorf("GetSuitableStuff() gotNotReservedProducers = %v, want %v", gotNotReservedProducers, tt.wantNotReservedProducers)
//			}
//			if !reflect.DeepEqual(gotNotReservedInstrumentalists, tt.wantNotReservedInstrumentalists) {
//				t.Errorf("GetSuitableStuff() gotNotReservedInstrumentalists = %v, want %v", gotNotReservedInstrumentalists, tt.wantNotReservedInstrumentalists)
//			}
//		})
//	}
//}
//
//func TestValidateTimeService_getNotReservedEquipments(t *testing.T) {
//	type fields struct {
//		roomRepo              repoInterface.IRoomRepository
//		equipmentRepo         repoInterface.IEquipmentRepository
//		producerRepo          repoInterface.IProducerRepository
//		instrumentalistRepo   repoInterface.IInstrumentalistRepository
//		reserveRepo           repoInterface.IReserveRepository
//		reservedEquipmentRepo repoInterface.IReservedEquipmentRepository
//	}
//
//	type args struct {
//		ctx     context.Context
//		request *dto.GetNotReservedEquipmentsRequest
//	}
//
//	tests := []struct {
//		name                      string
//		fields                    fields
//		args                      args
//		wantNotReservedEquipments [][]*model.Equipment
//		wantErr                   bool
//	}{
//		{
//			name: "test_pos_01",
//			args: args{
//				ctx: context.Background(),
//				request: &dto.GetNotReservedEquipmentsRequest{
//					ChoosenInterval: &model.TimeInterval{
//						StartTime: time.Date(2024, 4, 8, 13, 00, 00, 00, time.UTC),
//						EndTime:   time.Date(2024, 4, 8, 15, 00, 00, 00, time.UTC),
//					},
//					StudioId: 1,
//				},
//			},
//			wantErr: false,
//			wantNotReservedEquipments: [][]*model.Equipment{
//				{
//					&model.Equipment{
//						Id:            11,
//						Name:          "11",
//						EquipmentType: 1,
//						StudioId:      1,
//					},
//					&model.Equipment{
//						Id:            33,
//						Name:          "33",
//						EquipmentType: 1,
//						StudioId:      1,
//					},
//				},
//				{
//					&model.Equipment{
//						Id:            22,
//						Name:          "22",
//						EquipmentType: 2,
//						StudioId:      1,
//					},
//				},
//				{
//					&model.Equipment{
//						Id:            44,
//						Name:          "44",
//						EquipmentType: 3,
//						StudioId:      1,
//					},
//				},
//			},
//		},
//
//		{
//			name: "test_pos_02",
//			args: args{
//				ctx: context.Background(),
//				request: &dto.GetNotReservedEquipmentsRequest{
//					ChoosenInterval: &model.TimeInterval{
//						StartTime: time.Date(2024, 4, 8, 13, 00, 00, 00, time.UTC),
//						EndTime:   time.Date(2024, 4, 8, 15, 00, 00, 00, time.UTC),
//					},
//					StudioId: 2,
//				},
//			},
//			wantErr: false,
//			wantNotReservedEquipments: [][]*model.Equipment{
//				{
//					&model.Equipment{
//						Id:            11,
//						Name:          "11",
//						EquipmentType: 1,
//						StudioId:      2,
//					},
//					&model.Equipment{
//						Id:            33,
//						Name:          "33",
//						EquipmentType: 1,
//						StudioId:      2,
//					},
//				},
//				{
//					&model.Equipment{
//						Id:            55,
//						Name:          "55",
//						EquipmentType: 2,
//						StudioId:      2,
//					},
//				},
//				{
//					&model.Equipment{
//						Id:            44,
//						Name:          "44",
//						EquipmentType: 3,
//						StudioId:      2,
//					},
//				},
//			},
//		},
//	}
//
//	equipmentRepo := new(mocks.IEquipmentRepository)
//
//	for _, tt := range tests {
//
//		// test_pos_01
//
//		equipmentRepo.On("GetNotFullTimeFreeByStudioAndType",
//			//tt.args.request,
//			context.Background(),
//			&dto.GetEquipmentNotFullTimeFreeByStudioAndTypeRequest{
//				StudioId:     tt.args.request.StudioId,
//				Type:         1,
//				TimeInterval: tt.args.request.ChoosenInterval,
//			}).Return([]*dto.EquipmentAndTime{
//			{
//				&model.Equipment{
//					Id:            11,
//					Name:          "11",
//					EquipmentType: 1,
//					StudioId:      1,
//				},
//				&model.TimeInterval{
//					StartTime: time.Date(2024, 4, 8, 12, 00, 00, 00, time.UTC),
//					EndTime:   time.Date(2024, 4, 8, 13, 00, 00, 00, time.UTC),
//				},
//			},
//			{
//				&model.Equipment{
//					Id:            33,
//					Name:          "33",
//					EquipmentType: 1,
//					StudioId:      1,
//				},
//				&model.TimeInterval{
//					StartTime: time.Date(2024, 4, 8, 15, 00, 00, 00, time.UTC),
//					EndTime:   time.Date(2024, 4, 8, 17, 00, 00, 00, time.UTC),
//				},
//			},
//		}, nil)
//
//		equipmentRepo.On("GetNotFullTimeFreeByStudioAndType",
//			//tt.args.request,
//			context.Background(),
//			&dto.GetEquipmentNotFullTimeFreeByStudioAndTypeRequest{
//				StudioId:     tt.args.request.StudioId,
//				Type:         2,
//				TimeInterval: tt.args.request.ChoosenInterval,
//			}).Return(nil, nil)
//
//		equipmentRepo.On("GetNotFullTimeFreeByStudioAndType",
//			//tt.args.request,
//			context.Background(),
//			&dto.GetEquipmentNotFullTimeFreeByStudioAndTypeRequest{
//				StudioId:     tt.args.request.StudioId,
//				Type:         3,
//				TimeInterval: tt.args.request.ChoosenInterval,
//			}).Return(nil, nil)
//
//		equipmentRepo.On("GetFullTimeFreeByStudioAndType",
//			context.Background(),
//			&dto.GetEquipmentFullTimeFreeByStudioAndTypeRequest{
//				StudioId: tt.args.request.StudioId,
//				Type:     1,
//			}).Return(nil, nil)
//
//		equipmentRepo.On("GetFullTimeFreeByStudioAndType",
//			context.Background(),
//			&dto.GetEquipmentFullTimeFreeByStudioAndTypeRequest{
//				StudioId: tt.args.request.StudioId,
//				Type:     2,
//			}).Return([]*model.Equipment{
//			&model.Equipment{
//				Id:            22,
//				Name:          "22",
//				EquipmentType: 2,
//				StudioId:      1,
//			},
//		}, nil)
//
//		equipmentRepo.On("GetFullTimeFreeByStudioAndType",
//			context.Background(),
//			&dto.GetEquipmentFullTimeFreeByStudioAndTypeRequest{
//				StudioId: tt.args.request.StudioId,
//				Type:     3,
//			}).Return([]*model.Equipment{
//			{
//				Id:            44,
//				Name:          "44",
//				EquipmentType: 3,
//				StudioId:      1,
//			},
//		}, nil)
//
//		// test_pos_02
//
//		equipmentRepo.On("GetNotFullTimeFreeByStudioAndType",
//			//tt.args.request,
//			context.Background(),
//			&dto.GetEquipmentNotFullTimeFreeByStudioAndTypeRequest{
//				StudioId:     2,
//				Type:         1,
//				TimeInterval: tt.args.request.ChoosenInterval,
//			}).Return([]*dto.EquipmentAndTime{
//			{
//				&model.Equipment{
//					Id:            11,
//					Name:          "11",
//					EquipmentType: 1,
//					StudioId:      2,
//				},
//				&model.TimeInterval{
//					StartTime: time.Date(2024, 4, 8, 12, 00, 00, 00, time.UTC),
//					EndTime:   time.Date(2024, 4, 8, 13, 00, 00, 00, time.UTC),
//				},
//			},
//			{
//				&model.Equipment{
//					Id:            33,
//					Name:          "33",
//					EquipmentType: 1,
//					StudioId:      2,
//				},
//				&model.TimeInterval{
//					StartTime: time.Date(2024, 4, 8, 12, 00, 00, 00, time.UTC),
//					EndTime:   time.Date(2024, 4, 8, 13, 00, 00, 00, time.UTC),
//				},
//			},
//		}, nil)
//
//		equipmentRepo.On("GetFullTimeFreeByStudioAndType",
//			context.Background(),
//			&dto.GetEquipmentFullTimeFreeByStudioAndTypeRequest{
//				StudioId: 2,
//				Type:     1,
//			}).Return(nil, nil)
//
//		equipmentRepo.On("GetNotFullTimeFreeByStudioAndType",
//			//tt.args.request,
//			context.Background(),
//			&dto.GetEquipmentNotFullTimeFreeByStudioAndTypeRequest{
//				StudioId:     2,
//				Type:         2,
//				TimeInterval: tt.args.request.ChoosenInterval,
//			}).Return(nil, nil)
//
//		equipmentRepo.On("GetFullTimeFreeByStudioAndType",
//			context.Background(),
//			&dto.GetEquipmentFullTimeFreeByStudioAndTypeRequest{
//				StudioId: 2,
//				Type:     2,
//			}).Return([]*model.Equipment{
//			{
//				Id:            55,
//				Name:          "55",
//				EquipmentType: 2,
//				StudioId:      2,
//			},
//		}, nil)
//
//		equipmentRepo.On("GetNotFullTimeFreeByStudioAndType",
//			//tt.args.request,
//			context.Background(),
//			&dto.GetEquipmentNotFullTimeFreeByStudioAndTypeRequest{
//				StudioId:     2,
//				Type:         3,
//				TimeInterval: tt.args.request.ChoosenInterval,
//			}).Return(nil, nil)
//
//		equipmentRepo.On("GetFullTimeFreeByStudioAndType",
//			context.Background(),
//			&dto.GetEquipmentFullTimeFreeByStudioAndTypeRequest{
//				StudioId: 2,
//				Type:     3,
//			}).Return([]*model.Equipment{
//			{
//				Id:            44,
//				Name:          "44",
//				EquipmentType: 3,
//				StudioId:      2,
//			},
//		}, nil)
//
//		t.Run(tt.name, func(t *testing.T) {
//
//			s := ValidateTimeService{
//				//roomRepo:              tt.fields.roomRepo,
//				equipmentRepo: equipmentRepo,
//				//producerRepo:          tt.fields.producerRepo,
//				//instrumentalistRepo:   tt.fields.instrumentalistRepo,
//				//reserveRepo:           tt.fields.reserveRepo,
//				//reservedEquipmentRepo: tt.fields.reservedEquipmentRepo,
//			}
//			gotNotReservedEquipments, err := s.getNotReservedEquipments(tt.args.ctx, tt.args.request)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("getNotReservedEquipments() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(gotNotReservedEquipments, tt.wantNotReservedEquipments) {
//				t.Errorf("getNotReservedEquipments() gotNotReservedEquipments = %v, want %v", gotNotReservedEquipments, tt.wantNotReservedEquipments)
//			}
//		})
//	}
//}
//
//func TestValidateTimeService_getNotReservedInstrumentalists(t *testing.T) {
//	type fields struct {
//		roomRepo              repoInterface.IRoomRepository
//		equipmentRepo         repoInterface.IEquipmentRepository
//		producerRepo          repoInterface.IProducerRepository
//		instrumentalistRepo   repoInterface.IInstrumentalistRepository
//		reserveRepo           repoInterface.IReserveRepository
//		reservedEquipmentRepo repoInterface.IReservedEquipmentRepository
//	}
//	type args struct {
//		ctx     context.Context
//		request *dto.GetNotReservedInstrumentalistsRequest
//	}
//	tests := []struct {
//		name                            string
//		fields                          fields
//		args                            args
//		wantNotReservedInstrumentalists []*model.Instrumentalist
//		wantErr                         bool
//	}{
//		{
//			name: "test_pos_01",
//			args: args{
//				ctx: context.Background(),
//				request: &dto.GetNotReservedInstrumentalistsRequest{
//					ChoosenInterval: &model.TimeInterval{
//						StartTime: time.Date(2024, 4, 8, 15, 00, 00, 00, time.UTC),
//						EndTime:   time.Date(2024, 4, 8, 17, 00, 00, 00, time.UTC),
//					},
//					StudioId: 1,
//				},
//			},
//			wantNotReservedInstrumentalists: []*model.Instrumentalist{
//				&model.Instrumentalist{
//					Id:        1,
//					Name:      "1",
//					StudioId:  1,
//					StartHour: 9,
//					EndHour:   21,
//				},
//			},
//		},
//		{
//			name: "test_pos_02",
//			args: args{
//				ctx: context.Background(),
//				request: &dto.GetNotReservedInstrumentalistsRequest{
//					ChoosenInterval: &model.TimeInterval{
//						StartTime: time.Date(2024, 4, 8, 15, 00, 00, 00, time.UTC),
//						EndTime:   time.Date(2024, 4, 8, 17, 00, 00, 00, time.UTC),
//					},
//
//					StudioId: 2,
//				},
//			},
//			wantNotReservedInstrumentalists: []*model.Instrumentalist{
//				&model.Instrumentalist{
//					Id:        3,
//					Name:      "3",
//					StudioId:  2,
//					StartHour: 9,
//					EndHour:   21,
//				},
//			},
//		},
//		{
//			name: "test_pos_03",
//			args: args{
//				ctx: context.Background(),
//				request: &dto.GetNotReservedInstrumentalistsRequest{
//					ChoosenInterval: &model.TimeInterval{
//						StartTime: time.Date(2024, 4, 8, 13, 00, 00, 00, time.UTC),
//						EndTime:   time.Date(2024, 4, 8, 15, 00, 00, 00, time.UTC),
//					},
//
//					StudioId: 3,
//				},
//			},
//			wantNotReservedInstrumentalists: []*model.Instrumentalist{
//				&model.Instrumentalist{
//					Id:        31,
//					Name:      "1",
//					StudioId:  3,
//					StartHour: 9,
//					EndHour:   21,
//				},
//				&model.Instrumentalist{
//					Id:        32,
//					Name:      "1",
//					StudioId:  3,
//					StartHour: 13,
//					EndHour:   21,
//				},
//			},
//		},
//		{
//			name: "test_pos_04",
//			args: args{
//				ctx: context.Background(),
//				request: &dto.GetNotReservedInstrumentalistsRequest{
//					ChoosenInterval: &model.TimeInterval{
//						StartTime: time.Date(2024, 4, 8, 13, 00, 00, 00, time.UTC),
//						EndTime:   time.Date(2024, 4, 8, 15, 00, 00, 00, time.UTC),
//					},
//
//					StudioId: 4,
//				},
//			},
//			wantNotReservedInstrumentalists: nil,
//		},
//		{
//			name: "test_pos_05",
//			args: args{
//				ctx: context.Background(),
//				request: &dto.GetNotReservedInstrumentalistsRequest{
//					ChoosenInterval: &model.TimeInterval{
//						StartTime: time.Date(2024, 4, 8, 12, 00, 00, 00, time.UTC),
//						EndTime:   time.Date(2024, 4, 8, 15, 00, 00, 00, time.UTC),
//					},
//
//					StudioId: 5,
//				},
//			},
//			wantNotReservedInstrumentalists: nil,
//		},
//		//{
//		//	name: "test_neg_01",
//		//	args: args{
//		//		ctx: context.Background(),
//		//		request: &dto.GetNotReservedInstrumentalistsRequest{
//		//			ChoosenInterval: &model.TimeInterval{
//		//				StartTime: time.Date(2024, 4, 8, 12, 00, 00, 00, time.UTC),
//		//				EndTime:   time.Date(2024, 4, 8, 12, 00, 00, 00, time.UTC),
//		//			},
//		//			StudioId: 6,
//		//		},
//		//	},
//		//	wantErr:                         true,
//		//	wantNotReservedInstrumentalists: nil,
//		//},
//		{
//			name: "test_neg_02",
//			args: args{
//				ctx: context.Background(),
//				request: &dto.GetNotReservedInstrumentalistsRequest{
//					ChoosenInterval: &model.TimeInterval{
//						StartTime: time.Date(2024, 4, 8, 12, 00, 00, 00, time.UTC),
//						EndTime:   time.Date(2024, 4, 8, 12, 00, 00, 00, time.UTC),
//					},
//					StudioId: 0,
//				},
//			},
//			wantErr:                         true,
//			wantNotReservedInstrumentalists: nil,
//		},
//	}
//	prodRepo := new(mocks.IInstrumentalistRepository)
//	reserveRepo := new(mocks.IReserveRepository)
//
//	// test_pos_01
//
//	prodRepo.On("GetByStudio", context.Background(), &dto.GetInstrumentalistByStudioRequest{
//		StudioId: 1,
//	}).Return(
//		[]*model.Instrumentalist{
//			&model.Instrumentalist{
//				Id:        1,
//				Name:      "1",
//				StudioId:  1,
//				StartHour: 9,
//				EndHour:   21,
//			},
//		}, nil)
//
//	reserveRepo.On("GetByInstrumentalistId", context.Background(), &dto.GetReserveByInstrumentalistIdRequest{
//		InstrumentalistId: 1,
//	}).Return([]*model.Reserve{
//		&model.Reserve{
//			Id:                1,
//			UserId:            1,
//			RoomId:            1,
//			ProducerId:        1,
//			InstrumentalistId: 1,
//			TimeInterval: &model.TimeInterval{
//				StartTime: time.Date(2024, 4, 8, 12, 00, 00, 00, time.UTC),
//				EndTime:   time.Date(2024, 4, 8, 13, 00, 00, 00, time.UTC),
//			},
//		},
//	}, nil)
//
//	//test_pos_02
//
//	prodRepo.On("GetByStudio", context.Background(), &dto.GetInstrumentalistByStudioRequest{
//		StudioId: 2,
//	}).Return(
//		[]*model.Instrumentalist{
//			&model.Instrumentalist{
//				Id:        2,
//				Name:      "2",
//				StudioId:  2,
//				StartHour: 9,
//				EndHour:   15,
//			},
//			&model.Instrumentalist{
//				Id:        3,
//				Name:      "3",
//				StudioId:  2,
//				StartHour: 9,
//				EndHour:   21,
//			},
//		}, nil)
//
//	reserveRepo.On("GetByInstrumentalistId", context.Background(), &dto.GetReserveByInstrumentalistIdRequest{
//		InstrumentalistId: 2,
//	}).Return([]*model.Reserve{
//		&model.Reserve{
//			Id:                1,
//			UserId:            1,
//			RoomId:            1,
//			ProducerId:        1,
//			InstrumentalistId: 2,
//			TimeInterval: &model.TimeInterval{
//				StartTime: time.Date(2024, 4, 9, 12, 00, 00, 00, time.UTC),
//				EndTime:   time.Date(2024, 4, 9, 15, 00, 00, 00, time.UTC),
//			},
//		},
//	}, nil)
//
//	reserveRepo.On("GetByInstrumentalistId", context.Background(), &dto.GetReserveByInstrumentalistIdRequest{
//		InstrumentalistId: 3,
//	}).Return([]*model.Reserve{
//		&model.Reserve{
//			Id:                1,
//			UserId:            1,
//			RoomId:            1,
//			ProducerId:        1,
//			InstrumentalistId: 3,
//			TimeInterval: &model.TimeInterval{
//				StartTime: time.Date(2024, 4, 8, 12, 00, 00, 00, time.UTC),
//				EndTime:   time.Date(2024, 4, 8, 15, 00, 00, 00, time.UTC),
//			},
//		},
//	}, nil)
//
//	// test_pos_03
//
//	prodRepo.On("GetByStudio", context.Background(), &dto.GetInstrumentalistByStudioRequest{
//		StudioId: 3,
//	}).Return(
//		[]*model.Instrumentalist{
//			&model.Instrumentalist{
//				Id:        33,
//				Name:      "3",
//				StudioId:  3,
//				StartHour: 9,
//				EndHour:   14,
//			},
//			&model.Instrumentalist{
//				Id:        31,
//				Name:      "1",
//				StudioId:  3,
//				StartHour: 9,
//				EndHour:   21,
//			},
//			&model.Instrumentalist{
//				Id:        32,
//				Name:      "1",
//				StudioId:  3,
//				StartHour: 13,
//				EndHour:   21,
//			},
//			&model.Instrumentalist{
//				Id:        34,
//				Name:      "1",
//				StudioId:  3,
//				StartHour: 9,
//				EndHour:   21,
//			},
//		}, nil)
//
//	reserveRepo.On("GetByInstrumentalistId", context.Background(), &dto.GetReserveByInstrumentalistIdRequest{
//		InstrumentalistId: 33,
//	}).Return([]*model.Reserve{}, nil)
//
//	reserveRepo.On("GetByInstrumentalistId", context.Background(), &dto.GetReserveByInstrumentalistIdRequest{
//		InstrumentalistId: 31,
//	}).Return([]*model.Reserve{
//		&model.Reserve{
//			Id:                1,
//			UserId:            1,
//			RoomId:            1,
//			ProducerId:        1,
//			InstrumentalistId: 31,
//			TimeInterval: &model.TimeInterval{
//				StartTime: time.Date(2024, 4, 8, 12, 00, 00, 00, time.UTC),
//				EndTime:   time.Date(2024, 4, 8, 13, 00, 00, 00, time.UTC),
//			},
//		},
//		&model.Reserve{
//			Id:                1,
//			UserId:            1,
//			RoomId:            1,
//			ProducerId:        1,
//			InstrumentalistId: 31,
//			TimeInterval: &model.TimeInterval{
//				StartTime: time.Date(2024, 4, 8, 15, 00, 00, 00, time.UTC),
//				EndTime:   time.Date(2024, 4, 8, 17, 00, 00, 00, time.UTC),
//			},
//		},
//	}, nil)
//
//	reserveRepo.On("GetByInstrumentalistId", context.Background(), &dto.GetReserveByInstrumentalistIdRequest{
//		InstrumentalistId: 32,
//	}).Return([]*model.Reserve{
//		&model.Reserve{
//			Id:                1,
//			UserId:            1,
//			RoomId:            1,
//			ProducerId:        1,
//			InstrumentalistId: 32,
//			TimeInterval: &model.TimeInterval{
//				StartTime: time.Date(2024, 4, 8, 15, 00, 00, 00, time.UTC),
//				EndTime:   time.Date(2024, 4, 8, 17, 00, 00, 00, time.UTC),
//			},
//		},
//	}, nil)
//	reserveRepo.On("GetByInstrumentalistId", context.Background(), &dto.GetReserveByInstrumentalistIdRequest{
//		InstrumentalistId: 34,
//	}).Return([]*model.Reserve{
//		&model.Reserve{
//			Id:                1,
//			UserId:            1,
//			RoomId:            1,
//			ProducerId:        1,
//			InstrumentalistId: 34,
//			TimeInterval: &model.TimeInterval{
//				StartTime: time.Date(2024, 4, 8, 12, 00, 00, 00, time.UTC),
//				EndTime:   time.Date(2024, 4, 8, 17, 00, 00, 00, time.UTC),
//			},
//		},
//	}, nil)
//
//	// test_pos_04
//
//	prodRepo.On("GetByStudio", context.Background(), &dto.GetInstrumentalistByStudioRequest{
//		StudioId: 4,
//	}).Return(
//		[]*model.Instrumentalist{
//			&model.Instrumentalist{
//				Id:        43,
//				Name:      "3",
//				StudioId:  4,
//				StartHour: 9,
//				EndHour:   14,
//			},
//			&model.Instrumentalist{
//				Id:        41,
//				Name:      "1",
//				StudioId:  4,
//				StartHour: 9,
//				EndHour:   21,
//			},
//		}, nil)
//
//	reserveRepo.On("GetByInstrumentalistId", context.Background(), &dto.GetReserveByInstrumentalistIdRequest{
//		InstrumentalistId: 43,
//	}).Return([]*model.Reserve{}, nil)
//
//	reserveRepo.On("GetByInstrumentalistId", context.Background(), &dto.GetReserveByInstrumentalistIdRequest{
//		InstrumentalistId: 41,
//	}).Return([]*model.Reserve{
//		&model.Reserve{
//			Id:                1,
//			UserId:            1,
//			RoomId:            1,
//			ProducerId:        1,
//			InstrumentalistId: 41,
//			TimeInterval: &model.TimeInterval{
//				StartTime: time.Date(2024, 4, 8, 12, 00, 00, 00, time.UTC),
//				EndTime:   time.Date(2024, 4, 8, 17, 00, 00, 00, time.UTC),
//			},
//		},
//	}, nil)
//
//	//test_pos_05
//
//	prodRepo.On("GetByStudio", context.Background(), &dto.GetInstrumentalistByStudioRequest{
//		StudioId: 5,
//	}).Return(
//		[]*model.Instrumentalist{
//			&model.Instrumentalist{
//				Id:        53,
//				Name:      "3",
//				StudioId:  5,
//				StartHour: 9,
//				EndHour:   14,
//			},
//			&model.Instrumentalist{
//				Id:        51,
//				Name:      "1",
//				StudioId:  5,
//				StartHour: 9,
//				EndHour:   21,
//			},
//		}, nil)
//
//	reserveRepo.On("GetByInstrumentalistId", context.Background(), &dto.GetReserveByInstrumentalistIdRequest{
//		InstrumentalistId: 53,
//	}).Return([]*model.Reserve{}, nil)
//
//	reserveRepo.On("GetByInstrumentalistId", context.Background(), &dto.GetReserveByInstrumentalistIdRequest{
//		InstrumentalistId: 51,
//	}).Return([]*model.Reserve{
//		&model.Reserve{
//			Id:                1,
//			UserId:            1,
//			RoomId:            1,
//			ProducerId:        1,
//			InstrumentalistId: 51,
//			TimeInterval: &model.TimeInterval{
//				StartTime: time.Date(2024, 4, 8, 11, 00, 00, 00, time.UTC),
//				EndTime:   time.Date(2024, 4, 8, 13, 00, 00, 00, time.UTC),
//			},
//		},
//	}, nil)
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			s := ValidateTimeService{
//				roomRepo:              tt.fields.roomRepo,
//				equipmentRepo:         tt.fields.equipmentRepo,
//				producerRepo:          tt.fields.producerRepo,
//				instrumentalistRepo:   prodRepo,
//				reserveRepo:           reserveRepo,
//				reservedEquipmentRepo: tt.fields.reservedEquipmentRepo,
//			}
//			gotNotReservedInstrumentalists, err := s.getNotReservedInstrumentalists(tt.args.ctx, tt.args.request)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("getNotReservedInstrumentalists() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(gotNotReservedInstrumentalists, tt.wantNotReservedInstrumentalists) {
//				t.Errorf("getNotReservedInstrumentalists() gotNotReservedInstrumentalists = %v, want %v", gotNotReservedInstrumentalists, tt.wantNotReservedInstrumentalists)
//			}
//		})
//	}
//}
//
//func TestValidateTimeService_getNotReservedProducers(t *testing.T) {
//	type fields struct {
//		roomRepo              repoInterface.IRoomRepository
//		equipmentRepo         repoInterface.IEquipmentRepository
//		producerRepo          repoInterface.IProducerRepository
//		instrumentalistRepo   repoInterface.IInstrumentalistRepository
//		reserveRepo           repoInterface.IReserveRepository
//		reservedEquipmentRepo repoInterface.IReservedEquipmentRepository
//	}
//	type args struct {
//		ctx     context.Context
//		request *dto.GetNotReservedProducersRequest
//	}
//	tests := []struct {
//		name                     string
//		fields                   fields
//		args                     args
//		wantNotReservedProducers []*model.Producer
//		wantErr                  bool
//	}{
//		{
//			name: "test_pos_01",
//			args: args{
//				ctx: context.Background(),
//				request: &dto.GetNotReservedProducersRequest{
//					ChoosenInterval: &model.TimeInterval{
//						StartTime: time.Date(2024, 4, 8, 15, 00, 00, 00, time.UTC),
//						EndTime:   time.Date(2024, 4, 8, 17, 00, 00, 00, time.UTC),
//					},
//					StudioId: 1,
//				},
//			},
//			wantNotReservedProducers: []*model.Producer{
//				&model.Producer{
//					Id:        1,
//					Name:      "1",
//					StudioId:  1,
//					StartHour: 9,
//					EndHour:   21,
//				},
//			},
//		},
//		{
//			name: "test_pos_02",
//			args: args{
//				ctx: context.Background(),
//				request: &dto.GetNotReservedProducersRequest{
//					ChoosenInterval: &model.TimeInterval{
//						StartTime: time.Date(2024, 4, 8, 15, 00, 00, 00, time.UTC),
//						EndTime:   time.Date(2024, 4, 8, 17, 00, 00, 00, time.UTC),
//					},
//					StudioId: 2,
//				},
//			},
//			wantNotReservedProducers: []*model.Producer{
//				&model.Producer{
//					Id:        3,
//					Name:      "3",
//					StudioId:  2,
//					StartHour: 9,
//					EndHour:   21,
//				},
//			},
//		},
//
//		{
//			name: "test_pos_03",
//			args: args{
//				ctx: context.Background(),
//				request: &dto.GetNotReservedProducersRequest{
//					ChoosenInterval: &model.TimeInterval{
//						StartTime: time.Date(2024, 4, 8, 13, 00, 00, 00, time.UTC),
//						EndTime:   time.Date(2024, 4, 8, 15, 00, 00, 00, time.UTC),
//					},
//					StudioId: 3,
//				},
//			},
//			wantNotReservedProducers: []*model.Producer{
//				&model.Producer{
//					Id:        31,
//					Name:      "1",
//					StudioId:  3,
//					StartHour: 9,
//					EndHour:   21,
//				},
//				&model.Producer{
//					Id:        32,
//					Name:      "1",
//					StudioId:  3,
//					StartHour: 13,
//					EndHour:   21,
//				},
//			},
//		},
//		{
//			name: "test_pos_04",
//			args: args{
//				ctx: context.Background(),
//				request: &dto.GetNotReservedProducersRequest{
//					ChoosenInterval: &model.TimeInterval{
//						StartTime: time.Date(2024, 4, 8, 13, 00, 00, 00, time.UTC),
//						EndTime:   time.Date(2024, 4, 8, 15, 00, 00, 00, time.UTC),
//					},
//					StudioId: 4,
//				},
//			},
//			wantNotReservedProducers: nil,
//		},
//		{
//			name: "test_pos_05",
//			args: args{
//				ctx: context.Background(),
//				request: &dto.GetNotReservedProducersRequest{
//					ChoosenInterval: &model.TimeInterval{
//						StartTime: time.Date(2024, 4, 8, 12, 00, 00, 00, time.UTC),
//						EndTime:   time.Date(2024, 4, 8, 15, 00, 00, 00, time.UTC),
//					},
//					StudioId: 5,
//				},
//			},
//			wantNotReservedProducers: nil,
//		},
//		//{
//		//	name: "test_neg_01",
//		//	args: args{
//		//		ctx: context.Background(),
//		//		request: &dto.GetNotReservedProducersRequest{
//		//			ChoosenInterval: &model.TimeInterval{
//		//				StartTime: time.Date(2024, 4, 8, 12, 00, 00, 00, time.UTC),
//		//				EndTime:   time.Date(2024, 4, 8, 12, 00, 00, 00, time.UTC),
//		//			},
//		//			StudioId: 6,
//		//		},
//		//	},
//		//	wantErr:                  true,
//		//	wantNotReservedProducers: nil,
//		//},
//		{
//			name: "test_neg_02",
//			args: args{
//				ctx: context.Background(),
//				request: &dto.GetNotReservedProducersRequest{
//					ChoosenInterval: &model.TimeInterval{
//						StartTime: time.Date(2024, 4, 8, 12, 00, 00, 00, time.UTC),
//						EndTime:   time.Date(2024, 4, 8, 12, 00, 00, 00, time.UTC),
//					},
//					StudioId: 0,
//				},
//			},
//			wantErr:                  true,
//			wantNotReservedProducers: nil,
//		},
//	}
//	prodRepo := new(mocks.IProducerRepository)
//	reserveRepo := new(mocks.IReserveRepository)
//
//	// test_pos_01
//
//	prodRepo.On("GetByStudio", context.Background(), &dto.GetProducerByStudioRequest{
//		StudioId: 1,
//	}).Return(
//		[]*model.Producer{
//			&model.Producer{
//				Id:        1,
//				Name:      "1",
//				StudioId:  1,
//				StartHour: 9,
//				EndHour:   21,
//			},
//		}, nil)
//
//	reserveRepo.On("GetByProducerId", context.Background(), &dto.GetReserveByProducerIdRequest{
//		ProducerId: 1,
//	}).Return([]*model.Reserve{
//		&model.Reserve{
//			Id:                1,
//			UserId:            1,
//			RoomId:            1,
//			ProducerId:        1,
//			InstrumentalistId: 1,
//			TimeInterval: &model.TimeInterval{
//				StartTime: time.Date(2024, 4, 8, 12, 00, 00, 00, time.UTC),
//				EndTime:   time.Date(2024, 4, 8, 13, 00, 00, 00, time.UTC),
//			},
//		},
//	}, nil)
//
//	//test_pos_02
//
//	prodRepo.On("GetByStudio", context.Background(), &dto.GetProducerByStudioRequest{
//		StudioId: 2,
//	}).Return(
//		[]*model.Producer{
//			&model.Producer{
//				Id:        2,
//				Name:      "2",
//				StudioId:  2,
//				StartHour: 9,
//				EndHour:   15,
//			},
//			&model.Producer{
//				Id:        3,
//				Name:      "3",
//				StudioId:  2,
//				StartHour: 9,
//				EndHour:   21,
//			},
//		}, nil)
//
//	reserveRepo.On("GetByProducerId", context.Background(), &dto.GetReserveByProducerIdRequest{
//		ProducerId: 2,
//	}).Return([]*model.Reserve{
//		&model.Reserve{
//			Id:                1,
//			UserId:            1,
//			RoomId:            1,
//			ProducerId:        2,
//			InstrumentalistId: 1,
//			TimeInterval: &model.TimeInterval{
//				StartTime: time.Date(2024, 4, 9, 12, 00, 00, 00, time.UTC),
//				EndTime:   time.Date(2024, 4, 9, 15, 00, 00, 00, time.UTC),
//			},
//		},
//	}, nil)
//
//	reserveRepo.On("GetByProducerId", context.Background(), &dto.GetReserveByProducerIdRequest{
//		ProducerId: 3,
//	}).Return([]*model.Reserve{
//		&model.Reserve{
//			Id:                1,
//			UserId:            1,
//			RoomId:            1,
//			ProducerId:        3,
//			InstrumentalistId: 1,
//			TimeInterval: &model.TimeInterval{
//				StartTime: time.Date(2024, 4, 8, 12, 00, 00, 00, time.UTC),
//				EndTime:   time.Date(2024, 4, 8, 15, 00, 00, 00, time.UTC),
//			},
//		},
//	}, nil)
//
//	// test_pos_03
//
//	prodRepo.On("GetByStudio", context.Background(), &dto.GetProducerByStudioRequest{
//		StudioId: 3,
//	}).Return(
//		[]*model.Producer{
//			&model.Producer{
//				Id:        33,
//				Name:      "3",
//				StudioId:  3,
//				StartHour: 9,
//				EndHour:   14,
//			},
//			&model.Producer{
//				Id:        31,
//				Name:      "1",
//				StudioId:  3,
//				StartHour: 9,
//				EndHour:   21,
//			},
//			&model.Producer{
//				Id:        32,
//				Name:      "1",
//				StudioId:  3,
//				StartHour: 13,
//				EndHour:   21,
//			},
//			&model.Producer{
//				Id:        34,
//				Name:      "1",
//				StudioId:  3,
//				StartHour: 9,
//				EndHour:   21,
//			},
//		}, nil)
//
//	reserveRepo.On("GetByProducerId", context.Background(), &dto.GetReserveByProducerIdRequest{
//		ProducerId: 33,
//	}).Return([]*model.Reserve{}, nil)
//
//	reserveRepo.On("GetByProducerId", context.Background(), &dto.GetReserveByProducerIdRequest{
//		ProducerId: 31,
//	}).Return([]*model.Reserve{
//		&model.Reserve{
//			Id:                1,
//			UserId:            1,
//			RoomId:            1,
//			ProducerId:        31,
//			InstrumentalistId: 1,
//			TimeInterval: &model.TimeInterval{
//				StartTime: time.Date(2024, 4, 8, 12, 00, 00, 00, time.UTC),
//				EndTime:   time.Date(2024, 4, 8, 13, 00, 00, 00, time.UTC),
//			},
//		},
//		&model.Reserve{
//			Id:                1,
//			UserId:            1,
//			RoomId:            1,
//			ProducerId:        31,
//			InstrumentalistId: 1,
//			TimeInterval: &model.TimeInterval{
//				StartTime: time.Date(2024, 4, 8, 15, 00, 00, 00, time.UTC),
//				EndTime:   time.Date(2024, 4, 8, 17, 00, 00, 00, time.UTC),
//			},
//		},
//	}, nil)
//
//	reserveRepo.On("GetByProducerId", context.Background(), &dto.GetReserveByProducerIdRequest{
//		ProducerId: 32,
//	}).Return([]*model.Reserve{
//		&model.Reserve{
//			Id:                1,
//			UserId:            1,
//			RoomId:            1,
//			ProducerId:        32,
//			InstrumentalistId: 1,
//			TimeInterval: &model.TimeInterval{
//				StartTime: time.Date(2024, 4, 8, 15, 00, 00, 00, time.UTC),
//				EndTime:   time.Date(2024, 4, 8, 17, 00, 00, 00, time.UTC),
//			},
//		},
//	}, nil)
//	reserveRepo.On("GetByProducerId", context.Background(), &dto.GetReserveByProducerIdRequest{
//		ProducerId: 34,
//	}).Return([]*model.Reserve{
//		&model.Reserve{
//			Id:                1,
//			UserId:            1,
//			RoomId:            1,
//			ProducerId:        34,
//			InstrumentalistId: 1,
//			TimeInterval: &model.TimeInterval{
//				StartTime: time.Date(2024, 4, 8, 12, 00, 00, 00, time.UTC),
//				EndTime:   time.Date(2024, 4, 8, 17, 00, 00, 00, time.UTC),
//			},
//		},
//	}, nil)
//
//	// test_pos_04
//
//	prodRepo.On("GetByStudio", context.Background(), &dto.GetProducerByStudioRequest{
//		StudioId: 4,
//	}).Return(
//		[]*model.Producer{
//			&model.Producer{
//				Id:        43,
//				Name:      "3",
//				StudioId:  4,
//				StartHour: 9,
//				EndHour:   14,
//			},
//			&model.Producer{
//				Id:        41,
//				Name:      "1",
//				StudioId:  4,
//				StartHour: 9,
//				EndHour:   21,
//			},
//		}, nil)
//
//	reserveRepo.On("GetByProducerId", context.Background(), &dto.GetReserveByProducerIdRequest{
//		ProducerId: 43,
//	}).Return([]*model.Reserve{}, nil)
//
//	reserveRepo.On("GetByProducerId", context.Background(), &dto.GetReserveByProducerIdRequest{
//		ProducerId: 41,
//	}).Return([]*model.Reserve{
//		&model.Reserve{
//			Id:                1,
//			UserId:            1,
//			RoomId:            1,
//			ProducerId:        41,
//			InstrumentalistId: 1,
//			TimeInterval: &model.TimeInterval{
//				StartTime: time.Date(2024, 4, 8, 12, 00, 00, 00, time.UTC),
//				EndTime:   time.Date(2024, 4, 8, 17, 00, 00, 00, time.UTC),
//			},
//		},
//	}, nil)
//
//	//test_pos_05
//
//	prodRepo.On("GetByStudio", context.Background(), &dto.GetProducerByStudioRequest{
//		StudioId: 5,
//	}).Return(
//		[]*model.Producer{
//			&model.Producer{
//				Id:        53,
//				Name:      "3",
//				StudioId:  5,
//				StartHour: 9,
//				EndHour:   14,
//			},
//			&model.Producer{
//				Id:        51,
//				Name:      "1",
//				StudioId:  5,
//				StartHour: 9,
//				EndHour:   21,
//			},
//		}, nil)
//
//	reserveRepo.On("GetByProducerId", context.Background(), &dto.GetReserveByProducerIdRequest{
//		ProducerId: 53,
//	}).Return([]*model.Reserve{}, nil)
//
//	reserveRepo.On("GetByProducerId", context.Background(), &dto.GetReserveByProducerIdRequest{
//		ProducerId: 51,
//	}).Return([]*model.Reserve{
//		&model.Reserve{
//			Id:                1,
//			UserId:            1,
//			RoomId:            1,
//			ProducerId:        51,
//			InstrumentalistId: 1,
//			TimeInterval: &model.TimeInterval{
//				StartTime: time.Date(2024, 4, 8, 11, 00, 00, 00, time.UTC),
//				EndTime:   time.Date(2024, 4, 8, 13, 00, 00, 00, time.UTC),
//			},
//		},
//	}, nil)
//
//	for _, tt := range tests {
//
//		t.Run(tt.name, func(t *testing.T) {
//			s := ValidateTimeService{
//				roomRepo:              tt.fields.roomRepo,
//				equipmentRepo:         tt.fields.equipmentRepo,
//				producerRepo:          prodRepo,
//				instrumentalistRepo:   tt.fields.instrumentalistRepo,
//				reserveRepo:           reserveRepo,
//				reservedEquipmentRepo: tt.fields.reservedEquipmentRepo,
//			}
//			gotNotReservedProducers, err := s.getNotReservedProducers(tt.args.ctx, tt.args.request)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("getNotReservedProducers() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(gotNotReservedProducers, tt.wantNotReservedProducers) {
//				t.Errorf("getNotReservedProducers() gotNotReservedProducers = %v, want %v", gotNotReservedProducers, tt.wantNotReservedProducers)
//			}
//		})
//	}
//}
//
//func TestValidateTimeService_getNotReservedRooms(t *testing.T) {
//	type fields struct {
//		roomRepo              repoInterface.IRoomRepository
//		equipmentRepo         repoInterface.IEquipmentRepository
//		producerRepo          repoInterface.IProducerRepository
//		instrumentalistRepo   repoInterface.IInstrumentalistRepository
//		reserveRepo           repoInterface.IReserveRepository
//		reservedEquipmentRepo repoInterface.IReservedEquipmentRepository
//	}
//	type args struct {
//		ctx     context.Context
//		request *dto.GetNotReservedRoomsRequest
//	}
//	tests := []struct {
//		name                 string
//		fields               fields
//		args                 args
//		wantNotReservedRooms []*model.Room
//		wantErr              bool
//	}{
//		{
//			name: "test_pos_01",
//			args: args{
//				ctx: context.Background(),
//				request: &dto.GetNotReservedRoomsRequest{
//					ChoosenInterval: &model.TimeInterval{
//						StartTime: time.Date(2024, 4, 8, 15, 00, 00, 00, time.UTC),
//						EndTime:   time.Date(2024, 4, 8, 17, 00, 00, 00, time.UTC),
//					},
//					StudioId: 1,
//				},
//			},
//			wantNotReservedRooms: []*model.Room{
//				&model.Room{
//					Id:        1,
//					Name:      "1",
//					StudioId:  1,
//					StartHour: 9,
//					EndHour:   21,
//				},
//			},
//		},
//		{
//			name: "test_pos_02",
//			args: args{
//				ctx: context.Background(),
//				request: &dto.GetNotReservedRoomsRequest{
//					ChoosenInterval: &model.TimeInterval{
//						StartTime: time.Date(2024, 4, 8, 15, 00, 00, 00, time.UTC),
//						EndTime:   time.Date(2024, 4, 8, 17, 00, 00, 00, time.UTC),
//					},
//					StudioId: 2,
//				},
//			},
//			wantNotReservedRooms: []*model.Room{
//				&model.Room{
//					Id:        3,
//					Name:      "3",
//					StudioId:  2,
//					StartHour: 9,
//					EndHour:   21,
//				},
//			},
//		},
//		{
//			name: "test_pos_03",
//			args: args{
//				ctx: context.Background(),
//				request: &dto.GetNotReservedRoomsRequest{
//					ChoosenInterval: &model.TimeInterval{
//						StartTime: time.Date(2024, 4, 8, 13, 00, 00, 00, time.UTC),
//						EndTime:   time.Date(2024, 4, 8, 15, 00, 00, 00, time.UTC),
//					},
//					StudioId: 3,
//				},
//			},
//			wantNotReservedRooms: []*model.Room{
//				&model.Room{
//					Id:        31,
//					Name:      "1",
//					StudioId:  3,
//					StartHour: 9,
//					EndHour:   21,
//				},
//				&model.Room{
//					Id:        32,
//					Name:      "1",
//					StudioId:  3,
//					StartHour: 13,
//					EndHour:   21,
//				},
//			},
//		},
//		{
//			name: "test_pos_04",
//			args: args{
//				ctx: context.Background(),
//				request: &dto.GetNotReservedRoomsRequest{
//					ChoosenInterval: &model.TimeInterval{
//						StartTime: time.Date(2024, 4, 8, 13, 00, 00, 00, time.UTC),
//						EndTime:   time.Date(2024, 4, 8, 15, 00, 00, 00, time.UTC),
//					},
//					StudioId: 4,
//				},
//			},
//			wantNotReservedRooms: nil,
//		},
//		{
//			name: "test_pos_05",
//			args: args{
//				ctx: context.Background(),
//				request: &dto.GetNotReservedRoomsRequest{
//					ChoosenInterval: &model.TimeInterval{
//						StartTime: time.Date(2024, 4, 8, 12, 00, 00, 00, time.UTC),
//						EndTime:   time.Date(2024, 4, 8, 15, 00, 00, 00, time.UTC),
//					},
//					StudioId: 5,
//				},
//			},
//			wantNotReservedRooms: nil,
//		},
//		//{
//		//	name: "test_neg_01",
//		//	args: args{
//		//		ctx: context.Background(),
//		//		request: &dto.GetNotReservedRoomsRequest{
//		//			ChoosenInterval: &model.TimeInterval{
//		//				StartTime: time.Date(2024, 4, 8, 12, 00, 00, 00, time.UTC),
//		//				EndTime:   time.Date(2024, 4, 8, 12, 00, 00, 00, time.UTC),
//		//			},
//		//			StudioId: 6,
//		//		},
//		//	},
//		//	wantErr: true,
//		//	//wantNotReservedRooms: nil,
//		//},
//		{
//			name: "test_neg_02",
//			args: args{
//				ctx: context.Background(),
//				request: &dto.GetNotReservedRoomsRequest{
//					ChoosenInterval: &model.TimeInterval{
//						StartTime: time.Date(2024, 4, 8, 12, 00, 00, 00, time.UTC),
//						EndTime:   time.Date(2024, 4, 8, 12, 00, 00, 00, time.UTC),
//					},
//					StudioId: 0,
//				},
//			},
//			wantErr:              true,
//			wantNotReservedRooms: nil,
//		},
//	}
//	prodRepo := new(mocks.IRoomRepository)
//	reserveRepo := new(mocks.IReserveRepository)
//
//	// test_pos_01
//
//	prodRepo.On("GetByStudio", context.Background(), &dto.GetRoomByStudioRequest{
//		StudioId: 1,
//	}).Return(
//		[]*model.Room{
//			&model.Room{
//				Id:        1,
//				Name:      "1",
//				StudioId:  1,
//				StartHour: 9,
//				EndHour:   21,
//			},
//		}, nil)
//
//	reserveRepo.On("GetByRoomId", context.Background(), &dto.GetReserveByRoomIdRequest{
//		RoomId: 1,
//	}).Return([]*model.Reserve{
//		&model.Reserve{
//			Id:                1,
//			UserId:            1,
//			RoomId:            1,
//			ProducerId:        1,
//			InstrumentalistId: 1,
//			TimeInterval: &model.TimeInterval{
//				StartTime: time.Date(2024, 4, 8, 12, 00, 00, 00, time.UTC),
//				EndTime:   time.Date(2024, 4, 8, 13, 00, 00, 00, time.UTC),
//			},
//		},
//	}, nil)
//
//	//test_pos_02
//
//	prodRepo.On("GetByStudio", context.Background(), &dto.GetRoomByStudioRequest{
//		StudioId: 2,
//	}).Return(
//		[]*model.Room{
//			&model.Room{
//				Id:        2,
//				Name:      "2",
//				StudioId:  2,
//				StartHour: 9,
//				EndHour:   15,
//			},
//			&model.Room{
//				Id:        3,
//				Name:      "3",
//				StudioId:  2,
//				StartHour: 9,
//				EndHour:   21,
//			},
//		}, nil)
//
//	reserveRepo.On("GetByRoomId", context.Background(), &dto.GetReserveByRoomIdRequest{
//		RoomId: 2,
//	}).Return([]*model.Reserve{
//		&model.Reserve{
//			Id:                1,
//			UserId:            1,
//			RoomId:            2,
//			ProducerId:        1,
//			InstrumentalistId: 1,
//			TimeInterval: &model.TimeInterval{
//				StartTime: time.Date(2024, 4, 9, 12, 00, 00, 00, time.UTC),
//				EndTime:   time.Date(2024, 4, 9, 15, 00, 00, 00, time.UTC),
//			},
//		},
//	}, nil)
//
//	reserveRepo.On("GetByRoomId", context.Background(), &dto.GetReserveByRoomIdRequest{
//		RoomId: 3,
//	}).Return([]*model.Reserve{
//		&model.Reserve{
//			Id:                1,
//			UserId:            1,
//			RoomId:            3,
//			ProducerId:        1,
//			InstrumentalistId: 1,
//			TimeInterval: &model.TimeInterval{
//				StartTime: time.Date(2024, 4, 8, 12, 00, 00, 00, time.UTC),
//				EndTime:   time.Date(2024, 4, 8, 15, 00, 00, 00, time.UTC),
//			},
//		},
//	}, nil)
//
//	// test_pos_03
//
//	prodRepo.On("GetByStudio", context.Background(), &dto.GetRoomByStudioRequest{
//		StudioId: 3,
//	}).Return(
//		[]*model.Room{
//			&model.Room{
//				Id:        33,
//				Name:      "3",
//				StudioId:  3,
//				StartHour: 9,
//				EndHour:   14,
//			},
//			&model.Room{
//				Id:        31,
//				Name:      "1",
//				StudioId:  3,
//				StartHour: 9,
//				EndHour:   21,
//			},
//			&model.Room{
//				Id:        32,
//				Name:      "1",
//				StudioId:  3,
//				StartHour: 13,
//				EndHour:   21,
//			},
//			&model.Room{
//				Id:        34,
//				Name:      "1",
//				StudioId:  3,
//				StartHour: 9,
//				EndHour:   21,
//			},
//		}, nil)
//
//	reserveRepo.On("GetByRoomId", context.Background(), &dto.GetReserveByRoomIdRequest{
//		RoomId: 33,
//	}).Return([]*model.Reserve{}, nil)
//
//	reserveRepo.On("GetByRoomId", context.Background(), &dto.GetReserveByRoomIdRequest{
//		RoomId: 31,
//	}).Return([]*model.Reserve{
//		&model.Reserve{
//			Id:                1,
//			UserId:            1,
//			RoomId:            31,
//			ProducerId:        1,
//			InstrumentalistId: 1,
//			TimeInterval: &model.TimeInterval{
//				StartTime: time.Date(2024, 4, 8, 12, 00, 00, 00, time.UTC),
//				EndTime:   time.Date(2024, 4, 8, 13, 00, 00, 00, time.UTC),
//			},
//		},
//		&model.Reserve{
//			Id:                1,
//			UserId:            1,
//			RoomId:            31,
//			ProducerId:        1,
//			InstrumentalistId: 1,
//			TimeInterval: &model.TimeInterval{
//				StartTime: time.Date(2024, 4, 8, 15, 00, 00, 00, time.UTC),
//				EndTime:   time.Date(2024, 4, 8, 17, 00, 00, 00, time.UTC),
//			},
//		},
//	}, nil)
//
//	reserveRepo.On("GetByRoomId", context.Background(), &dto.GetReserveByRoomIdRequest{
//		RoomId: 32,
//	}).Return([]*model.Reserve{
//		&model.Reserve{
//			Id:                1,
//			UserId:            1,
//			RoomId:            32,
//			ProducerId:        1,
//			InstrumentalistId: 1,
//			TimeInterval: &model.TimeInterval{
//				StartTime: time.Date(2024, 4, 8, 15, 00, 00, 00, time.UTC),
//				EndTime:   time.Date(2024, 4, 8, 17, 00, 00, 00, time.UTC),
//			},
//		},
//	}, nil)
//	reserveRepo.On("GetByRoomId", context.Background(), &dto.GetReserveByRoomIdRequest{
//		RoomId: 34,
//	}).Return([]*model.Reserve{
//		&model.Reserve{
//			Id:                1,
//			UserId:            1,
//			RoomId:            34,
//			ProducerId:        1,
//			InstrumentalistId: 1,
//			TimeInterval: &model.TimeInterval{
//				StartTime: time.Date(2024, 4, 8, 12, 00, 00, 00, time.UTC),
//				EndTime:   time.Date(2024, 4, 8, 17, 00, 00, 00, time.UTC),
//			},
//		},
//	}, nil)
//
//	// test_pos_04
//
//	prodRepo.On("GetByStudio", context.Background(), &dto.GetRoomByStudioRequest{
//		StudioId: 4,
//	}).Return(
//		[]*model.Room{
//			&model.Room{
//				Id:        43,
//				Name:      "3",
//				StudioId:  4,
//				StartHour: 9,
//				EndHour:   14,
//			},
//			&model.Room{
//				Id:        41,
//				Name:      "1",
//				StudioId:  4,
//				StartHour: 9,
//				EndHour:   21,
//			},
//		}, nil)
//
//	reserveRepo.On("GetByRoomId", context.Background(), &dto.GetReserveByRoomIdRequest{
//		RoomId: 43,
//	}).Return([]*model.Reserve{}, nil)
//
//	reserveRepo.On("GetByRoomId", context.Background(), &dto.GetReserveByRoomIdRequest{
//		RoomId: 41,
//	}).Return([]*model.Reserve{
//		&model.Reserve{
//			Id:                1,
//			UserId:            1,
//			RoomId:            41,
//			ProducerId:        1,
//			InstrumentalistId: 1,
//			TimeInterval: &model.TimeInterval{
//				StartTime: time.Date(2024, 4, 8, 12, 00, 00, 00, time.UTC),
//				EndTime:   time.Date(2024, 4, 8, 17, 00, 00, 00, time.UTC),
//			},
//		},
//	}, nil)
//
//	//test_pos_05
//
//	prodRepo.On("GetByStudio", context.Background(), &dto.GetRoomByStudioRequest{
//		StudioId: 5,
//	}).Return(
//		[]*model.Room{
//			&model.Room{
//				Id:        53,
//				Name:      "3",
//				StudioId:  5,
//				StartHour: 9,
//				EndHour:   14,
//			},
//			&model.Room{
//				Id:        51,
//				Name:      "1",
//				StudioId:  5,
//				StartHour: 9,
//				EndHour:   21,
//			},
//		}, nil)
//
//	reserveRepo.On("GetByRoomId", context.Background(), &dto.GetReserveByRoomIdRequest{
//		RoomId: 53,
//	}).Return([]*model.Reserve{}, nil)
//
//	reserveRepo.On("GetByRoomId", context.Background(), &dto.GetReserveByRoomIdRequest{
//		RoomId: 51,
//	}).Return([]*model.Reserve{
//		&model.Reserve{
//			Id:                1,
//			UserId:            1,
//			RoomId:            51,
//			ProducerId:        1,
//			InstrumentalistId: 1,
//			TimeInterval: &model.TimeInterval{
//				StartTime: time.Date(2024, 4, 8, 11, 00, 00, 00, time.UTC),
//				EndTime:   time.Date(2024, 4, 8, 13, 00, 00, 00, time.UTC),
//			},
//		},
//	}, nil)
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			s := ValidateTimeService{
//				roomRepo:              prodRepo,
//				equipmentRepo:         tt.fields.equipmentRepo,
//				producerRepo:          tt.fields.producerRepo,
//				instrumentalistRepo:   tt.fields.instrumentalistRepo,
//				reserveRepo:           reserveRepo,
//				reservedEquipmentRepo: tt.fields.reservedEquipmentRepo,
//			}
//			gotNotReservedRooms, err := s.getNotReservedRooms(tt.args.ctx, tt.args.request)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("getNotReservedRooms() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(gotNotReservedRooms, tt.wantNotReservedRooms) {
//				t.Errorf("getNotReservedRooms() gotNotReservedRooms = %v, want %v", gotNotReservedRooms, tt.wantNotReservedRooms)
//			}
//		})
//	}
//}
