package role

import (
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/role"
)

func Ent2Grpc(row *npool.Role) *npool.Role {
	if row == nil {
		return nil
	}
	row.Default = row.DefaultInt != 0
	return row
}

func Ent2GrpcMany(rows []*npool.Role) []*npool.Role {
	roles := []*npool.Role{}
	for _, row := range rows {
		roles = append(roles, Ent2Grpc(row))
	}
	return roles
}
