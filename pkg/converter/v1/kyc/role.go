package kyc

import (
	mgrkyc "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/kyc"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/kyc"
)

func Ent2Grpc(row *npool.Kyc) *npool.Kyc {
	if row == nil {
		return nil
	}
	row.DocumentType = mgrkyc.KycDocumentType(mgrkyc.KycDocumentType_value[row.GetDocumentTypeStr()])
	row.EntityType = mgrkyc.KycEntityType(mgrkyc.KycEntityType_value[row.GetEntityTypeStr()])
	return row
}

func Ent2GrpcMany(rows []*npool.Kyc) []*npool.Kyc {
	roles := []*npool.Kyc{}
	for _, row := range rows {
		roles = append(roles, Ent2Grpc(row))
	}
	return roles
}
