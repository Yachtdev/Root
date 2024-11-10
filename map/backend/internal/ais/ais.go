package ais

import (
	"github.com/BertoldVdb/go-ais"
	"github.com/BertoldVdb/go-ais/aisnmea"
	"go.uber.org/zap"
)

type AIS struct {
	nmea   *aisnmea.NMEACodec
	logger *zap.SugaredLogger
}

func New(logger *zap.SugaredLogger) *AIS {
	return &AIS{
		nmea:   aisnmea.NMEACodecNew(ais.CodecNew(false, false)),
		logger: logger,
	}
}

func (a *AIS) Encode(lat, lon, cog float64) []string {
	encoded := a.nmea.EncodeSentence(aisnmea.VdmPacket{
		MessageType: "AIVDM",
		Packet: ais.PositionReport{
			Longitude:        ais.FieldLatLonFine(lon),
			Latitude:         ais.FieldLatLonFine(lat),
			Valid:            true,
			PositionAccuracy: true,
			Raim:             false,
			Header: ais.Header{
				MessageID: 1,
			},
			Timestamp: 7,
			Cog:       ais.Field10(cog),
		},
	})

	return encoded
}

func (a *AIS) Decode(msg []byte) {}
