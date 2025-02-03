package misc

func (ms *MiscService) GetDatabaseStatus() string {
	status := ms.miscRepository.GetDatabaseStatus()

	if status {
		return DatabaseIsHealthy
	} else {
		return DatabaseIsOffline
	}
}
