package kyc

import (
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/kyc"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
)

func Ent2Grpc(row *npool.Kyc) *npool.Kyc {
	if row == nil {
		return nil
	}
	row.DocumentType = basetypes.KycDocumentType(basetypes.KycDocumentType_value[row.GetDocumentTypeStr()])
	row.EntityType = basetypes.KycEntityType(basetypes.KycEntityType_value[row.GetEntityTypeStr()])
	row.State = basetypes.KycState(basetypes.KycState_value[row.GetStateStr()])
	return row
}

func Ent2GrpcMany(rows []*npool.Kyc) []*npool.Kyc {
	roles := []*npool.Kyc{}
	for _, row := range rows {
		roles = append(roles, Ent2Grpc(row))
	}
	return roles
}
