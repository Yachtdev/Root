package ais

import (
	"github.com/BertoldVdb/go-ais"
	"github.com/BertoldVdb/go-ais/aisnmea"
	"go.uber.org/zap"
)

// AIS производит кодирование и декодирование строк AIS.
type (
	AIS struct {
		nmea   *aisnmea.NMEACodec
		logger *zap.SugaredLogger
	}
)

// New AIS-constructor.
func New(logger *zap.SugaredLogger) *AIS {
	return &AIS{
		nmea:   aisnmea.NMEACodecNew(ais.CodecNew(false, false)),
		logger: logger,
	}
}

// Encode кодирует в AIS.
func (a *AIS) Encode(lat, lon, cog float32) []string {
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

// Decode декодирует из AIS.
func (a *AIS) Decode(msg []byte) {}
