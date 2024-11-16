package service

// SendAisPointToMqtt производит кодирование строки в формат AIS и отправляет на mqtt-сервер.
func (s *Service) SendAisPointToMqtt(lat, lon, course float32) {
	strs := s.AIS.Encode(lat, lon, course)

	for _, str := range strs {
		s.MQTT.Pub([]byte(str), "test")
	}
}
