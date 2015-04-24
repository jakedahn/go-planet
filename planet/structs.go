package planet

type CoordType float64
type Coordinate [2]CoordType

type FeatureCollection struct {
	Features []Feature `json:"features"`
	Links    struct {
		Frist string `json:"first"`
		Next  string `json:"next"`
		Prev  string `json:"prev"`
		Self  string `json:"self"`
	} `json:"links"`
}

type Properties struct {
	Acquired string `json:"acquired,omitempty"`
	// FileSize int32  `json:"file_size,omitempty"`
	// // CloudCover struct {
	// // 	Estimated float64 `json:"estimated,omitempty"`
	// // } `json:"cloud_cover,omitempty"`

	// // ImageStatistics struct {
	// // 	SNR float64 `snr:"foo,omitempty"`
	// // } `json:"image_statistics,omitempty"`

	// // Links struct {
	// // 	Full            string `json:"full,omitempty"`
	// // 	Self            string `json:"self,omitempty"`
	// // 	SquareThumbnail string `json:"square_thumbnail,omitempty"`
	// // 	Thumbnail       string `json:"thumbnail,omitempty"`
	// // } `json:"links,omitempty"`

	// // Sat struct {
	// // 	Alt      float64 `json:"alt,omitempty"`
	// // 	Lat      float64 `json:"lat,omitempty"`
	// // 	Lng      float64 `json:"lng,omitempty"`
	// // 	OffNadir float64 `json:"off_nadir,omitempty"`
	// // } `json:"sat,omitempty"`

	// // Sun struct {
	// // 	Altitude       float64 `json:"altitude,omitempty"`
	// // 	Azimuth        float64 `json:"azimuth,omitempty"`
	// // 	LocalTimeOfDat float64 `json:"local_time_of_day,omitempty"`
	// // } `json:"sun,omitempty"`
}

type Geometry struct {
	Coordinates [][]Coordinate `json:"coordinates,omitempty"`
}

type Feature struct {
	Id         string     `json:"id"`
	Geometry   Geometry   `json:"geometry,omitempty"`
	Properties Properties `json:"properties,omitempty"`
}

type FeatureListParams struct {
	SatId                                 string `url:"sat.id,ommitempty"`
	CameraBitDepth                        string `url:"camera.bit_depth,ommitempty"`
	CameraColorMode                       string `url:"camera.color_mode,ommitempty"`
	Acquired                              string `url:"acquired,ommitempty"`
	AcquiredGreaterThan                   string `url:"acquired.gt,ommitempty"`
	AcquiredGreaterThanOrEqual            string `url:"acquired.gte,ommitempty"`
	AcquiredLessThan                      string `url:"acquired.lt,ommitempty"`
	AcquiredLessThanOrEqual               string `url:"acquired.lte,ommitempty"`
	CameraExposureTime                    string `url:"camera.exposure_time,ommitempty"`
	CameraExposureTimeGreaterThan         string `url:"camera.exposure_time.gt,ommitempty"`
	CameraExposureTimeGreaterThanOrEqual  string `url:"camera.exposure_time.gte,ommitempty"`
	CameraExposureTimeLessThan            string `url:"camera.exposure_time.lt,ommitempty"`
	CameraExposureTimeLessThanOrEqual     string `url:"camera.exposure_time.lte,ommitempty"`
	CameraGain                            string `url:"camera.gain,ommitempty"`
	CameraGainGreaterThan                 string `url:"camera.gain.gt,ommitempty"`
	CameraGainGreaterThanOrEqual          string `url:"camera.gain.gte,ommitempty"`
	CameraGainLessThan                    string `url:"camera.gain.lt,ommitempty"`
	CameraGainLessThanOrEqual             string `url:"camera.gain.lte,ommitempty"`
	CameraTDIPulses                       string `url:"camera.tdi_pulses,ommitempty"`
	CameraTDIPulsesGreaterThan            string `url:"camera.tdi_pulses.gt,ommitempty"`
	CameraTDIPulsesGreaterThanOrEqual     string `url:"camera.tdi_pulses.gte,ommitempty"`
	CameraTDIPulsesLessThan               string `url:"camera.tdi_pulses.lt,ommitempty"`
	CameraTDIPulsesLessThanOrEqual        string `url:"camera.tdi_pulses.lte,ommitempty"`
	CloudCoverEstimated                   string `url:"cloud_cover.estimated,ommitempty"`
	CloudCoverEstimatedGreaterThan        string `url:"cloud_cover.estimated.gt,ommitempty"`
	CloudCoverEstimatedGreaterThanOrEqual string `url:"cloud_cover.estimated.gte,ommitempty"`
	CloudCoverEstimatedLessThan           string `url:"cloud_cover.estimated.lt,ommitempty"`
	CloudCoverEstimatedLessThanOrEqual    string `url:"cloud_cover.estimated.lte,ommitempty"`
	FileSize                              string `url:"file_size,ommitempty"`
	FileSizeGreaterThan                   string `url:"file_size.gt,ommitempty"`
	FileSizeGreaterThanOrEqual            string `url:"file_size.gte,ommitempty"`
	FileSizeLessThan                      string `url:"file_size.lt,ommitempty"`
	FileSizeLessThanOrEqual               string `url:"file_size.lte,ommitempty"`
	ImageStatisticsGSD                    string `url:"image_statistics.gsd,ommitempty"`
	ImageStatisticsGSDGreaterThan         string `url:"image_statistics.gsd.gt,ommitempty"`
	ImageStatisticsGSDGreaterThanOrEqual  string `url:"image_statistics.gsd.gte,ommitempty"`
	ImageStatisticsGSDLessThan            string `url:"image_statistics.gsd.lt,ommitempty"`
	ImageStatisticsGSDLessThanOrEqual     string `url:"image_statistics.gsd.lte,ommitempty"`
	ImageStatisticsSNR                    string `url:"image_statistics.snr,ommitempty"`
	ImageStatisticsSNRGreaterThan         string `url:"image_statistics.snr.gt,ommitempty"`
	ImageStatisticsSNRGreaterThanOrEqual  string `url:"image_statistics.snr.gte,ommitempty"`
	ImageStatisticsSNRLessThan            string `url:"image_statistics.snr.lt,ommitempty"`
	ImageStatisticsSNRLessThanOrEqual     string `url:"image_statistics.snr.lte,ommitempty"`
	SatAltitude                           string `url:"sat.alt,ommitempty"`
	SatAltitudeGreaterThan                string `url:"sat.alt.gt,ommitempty"`
	SatAltitudeGreaterThanOrEqual         string `url:"sat.alt.gte,ommitempty"`
	SatAltitudeLessThan                   string `url:"sat.alt.lt,ommitempty"`
	SatAltitudeLessThanOrEqual            string `url:"sat.alt.lte,ommitempty"`
	SatLat                                string `url:"sat.lat,ommitempty"`
	SatLatGreaterThan                     string `url:"sat.lat.gt,ommitempty"`
	SatLatGreaterThanOrEqual              string `url:"sat.lat.gte,ommitempty"`
	SatLatLessThan                        string `url:"sat.lat.lt,ommitempty"`
	SatLatLessThanOrEqual                 string `url:"sat.lat.lte,ommitempty"`
	SatLng                                string `url:"sat.lng,ommitempty"`
	SatLngGreaterThan                     string `url:"sat.lng.gt,ommitempty"`
	SatLngGreaterThanOrEqual              string `url:"sat.lng.gte,ommitempty"`
	SatLngLessThan                        string `url:"sat.lng.lt,ommitempty"`
	SatLngLessThanOrEqual                 string `url:"sat.lng.lte,ommitempty"`
	SatOffNadir                           string `url:"sat.off_nadir,ommitempty"`
	SatOffNadirGreaterThan                string `url:"sat.off_nadir.gt,ommitempty"`
	SatOffNadirGreaterThanOrEqual         string `url:"sat.off_nadir.gte,ommitempty"`
	SatOffNadirLessThan                   string `url:"sat.off_nadir.lt,ommitempty"`
	SatOffNadirLessThanOrEqual            string `url:"sat.off_nadir.lte,ommitempty"`
	StripId                               string `url:"strip_id,ommitempty"`
	StripIdGreaterThan                    string `url:"strip_id.gt,ommitempty"`
	StripIdGreaterThanOrEqual             string `url:"strip_id.gte,ommitempty"`
	StripIdLessThan                       string `url:"strip_id.lt,ommitempty"`
	StripIdLessThanOrEqual                string `url:"strip_id.lte,ommitempty"`
	SunAltitude                           string `url:"sun.altitude,ommitempty"`
	SunAltitudeGreaterThan                string `url:"sun.altitude.gt,ommitempty"`
	SunAltitudeGreaterThanOrEqual         string `url:"sun.altitude.gte,ommitempty"`
	SunAltitudeLessThan                   string `url:"sun.altitude.lt,ommitempty"`
	SunAltitudeLessThanOrEqual            string `url:"sun.altitude.lte,ommitempty"`
	SunAzimuth                            string `url:"sun.azimuth,ommitempty"`
	SunAzimuthGreaterThan                 string `url:"sun.azimuth.gt,ommitempty"`
	SunAzimuthGreaterThanOrEqual          string `url:"sun.azimuth.gte,ommitempty"`
	SunAzimuthLessThan                    string `url:"sun.azimuth.lt,ommitempty"`
	SunAzimuthLessThanOrEqual             string `url:"sun.azimuth.lte,ommitempty"`
	SunLocalTimeOfDay                     string `url:"sun.local_time_of_day,ommitempty"`
	SunLocalTimeOfDayGreaterThan          string `url:"sun.local_time_of_day.gt,ommitempty"`
	SunLocalTimeOfDayGreaterThanOrEqual   string `url:"sun.local_time_of_day.gte,ommitempty"`
	SunLocalTimeOfDayLessThan             string `url:"sun.local_time_of_day.lt,ommitempty"`
	SunLocalTimeOfDayLessThanOrEqual      string `url:"sun.local_time_of_day.lte,ommitempty"`
}
