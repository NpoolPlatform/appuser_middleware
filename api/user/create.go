package user

//
//func (s *Server) CreateUser(ctx context.Context, in *npool.CreateUserRequest) (*npool.CreateUserResponse, error) {
//	if _, err := uuid.Parse(in.GetAppID()); err != nil {
//		logger.Sugar().Errorw("GetUser", "error", err)
//		return &npool.GetUserResponse{}, status.Error(codes.InvalidArgument, "AppID is invalid")
//	}
//	if _, err := uuid.Parse(in.GetUserID()); err != nil {
//		logger.Sugar().Errorw("GetUser", "error", err)
//		return &npool.GetUserResponse{}, status.Error(codes.InvalidArgument, "UserID is invalid")
//	}
//
//	info, err := muser.GetUser(ctx, in.GetAppID(), in.GetUserID())
//	if err != nil {
//		logger.Sugar().Errorw("GetUser", "error", err)
//		return &npool.GetUserResponse{}, status.Error(codes.Internal, "fail get user")
//	}
//
//	ginfo, err := cuser.QueryEnt2Grpc(info)
//	if err != nil {
//		logger.Sugar().Errorw("GetUser", "error", err)
//		return &npool.GetUserResponse{}, status.Error(codes.Internal, "invalid value")
//	}
//
//	return &npool.GetUserResponse{
//		Info: ginfo,
//	}, nil
//}
//
//func (s *Server) GetUsers(ctx context.Context, in *npool.GetUsersRequest) (*npool.GetUsersResponse, error) {
//	if _, err := uuid.Parse(in.GetAppID()); err != nil {
//		logger.Sugar().Errorw("GetUsers", "error", err)
//		return &npool.GetUsersResponse{}, status.Error(codes.InvalidArgument, "AppID is invalid")
//	}
//
//	infos, err := muser.GetUsers(ctx, in.GetAppID(), in.GetOffset(), in.GetLimit())
//	if err != nil {
//		logger.Sugar().Errorw("GetUsers", "error", err)
//		return &npool.GetUsersResponse{}, status.Error(codes.Internal, "fail get users")
//	}
//
//	ginfos := []*npool.User{}
//	for _, val := range infos {
//		ginfo, err := cuser.QueryEnt2Grpc(val)
//		if err != nil {
//			logger.Sugar().Errorw("GetUsers", "error", err)
//			return &npool.GetUsersResponse{}, status.Error(codes.Internal, "invalid value")
//		}
//		ginfos = append(ginfos, ginfo)
//	}
//
//	return &npool.GetUsersResponse{
//		Infos: ginfos,
//	}, nil
//}
//
//func (s *Server) GetManyUsers(ctx context.Context, in *npool.GetManyUsersRequest) (*npool.GetManyUsersResponse, error) {
//	if len(in.IDs) == 0{
//		logger.Sugar().Errorw("GetManyUsers", "ids empty")
//		return &npool.GetManyUsersResponse{}, status.Error(codes.InvalidArgument, "ids empty")
//	}
//
//	infos, err := muser.GetManyUsers(ctx, in.IDs)
//	if err != nil {
//		logger.Sugar().Errorw("GetManyUsers", "error", err)
//		return &npool.GetManyUsersResponse{}, status.Error(codes.Internal, "fail get many users")
//	}
//
//	ginfos := []*npool.User{}
//	for _, val := range infos {
//		ginfo, err := cuser.QueryEnt2Grpc(val)
//		if err != nil {
//			logger.Sugar().Errorw("GetManyUsers", "error", err)
//			return &npool.GetManyUsersResponse{}, status.Error(codes.Internal, "invalid value")
//		}
//		ginfos = append(ginfos, ginfo)
//	}
//
//	return &npool.GetManyUsersResponse{
//		Infos: ginfos,
//	}, nil
//}
