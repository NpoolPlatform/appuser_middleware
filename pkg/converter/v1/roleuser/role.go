package roleuser

import (
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/role"
)

func Ent2Grpc(row *npool.RoleUser) *npool.RoleUser {
	if row == nil {
		return nil
	}
	row.Default = row.DefaultInt != 0
	return row
}

func Ent2GrpcMany(rows []*npool.RoleUser) []*npool.RoleUser {
	roles := []*npool.RoleUser{}
	for _, row := range rows {
		roles = append(roles, Ent2Grpc(row))
	}
	return roles
}
