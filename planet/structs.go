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
	Acquired   string `json:"acquired,omitempty"`
	FileSize   int32  `json:"file_size,omitempty"`
	CloudCover struct {
		Estimated float64 `json:"estimated,omitempty"`
	} `json:"cloud_cover,omitempty"`

	ImageStatistics struct {
		SNR float64 `snr:"foo,omitempty"`
	} `json:"image_statistics,omitempty"`

	Links struct {
		Full            string `json:"full,omitempty"`
		Self            string `json:"self,omitempty"`
		SquareThumbnail string `json:"square_thumbnail,omitempty"`
		Thumbnail       string `json:"thumbnail,omitempty"`
	} `json:"links,omitempty"`

	Sat struct {
		Alt      float64 `json:"alt,omitempty"`
		Lat      float64 `json:"lat,omitempty"`
		Lng      float64 `json:"lng,omitempty"`
		OffNadir float64 `json:"off_nadir,omitempty"`
	} `json:"sat,omitempty"`

	Sun struct {
		Altitude       float64 `json:"altitude,omitempty"`
		Azimuth        float64 `json:"azimuth,omitempty"`
		LocalTimeOfDat float64 `json:"local_time_of_day,omitempty"`
	} `json:"sun,omitempty"`
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
	SatId                                 string `url:"sat.id,omitempty"`
	CameraBitDepth                        string `url:"camera.bit_depth,omitempty"`
	CameraColorMode                       string `url:"camera.color_mode,omitempty"`
	Acquired                              string `url:"acquired,omitempty"`
	AcquiredLaterThan                     string `url:"acquired.gt,omitempty"`
	AcquiredLaterThanOrEqual              string `url:"acquired.gte,omitempty"`
	AcquiredEarlierThan                   string `url:"acquired.lt,omitempty"`
	AcquiredEarlierThanOrEqual            string `url:"acquired.lte,omitempty"`
	CameraExposureTime                    string `url:"camera.exposure_time,omitempty"`
	CameraExposureTimeGreaterThan         string `url:"camera.exposure_time.gt,omitempty"`
	CameraExposureTimeGreaterThanOrEqual  string `url:"camera.exposure_time.gte,omitempty"`
	CameraExposureTimeLessThan            string `url:"camera.exposure_time.lt,omitempty"`
	CameraExposureTimeLessThanOrEqual     string `url:"camera.exposure_time.lte,omitempty"`
	CameraGain                            string `url:"camera.gain,omitempty"`
	CameraGainGreaterThan                 string `url:"camera.gain.gt,omitempty"`
	CameraGainGreaterThanOrEqual          string `url:"camera.gain.gte,omitempty"`
	CameraGainLessThan                    string `url:"camera.gain.lt,omitempty"`
	CameraGainLessThanOrEqual             string `url:"camera.gain.lte,omitempty"`
	CameraTDIPulses                       string `url:"camera.tdi_pulses,omitempty"`
	CameraTDIPulsesGreaterThan            string `url:"camera.tdi_pulses.gt,omitempty"`
	CameraTDIPulsesGreaterThanOrEqual     string `url:"camera.tdi_pulses.gte,omitempty"`
	CameraTDIPulsesLessThan               string `url:"camera.tdi_pulses.lt,omitempty"`
	CameraTDIPulsesLessThanOrEqual        string `url:"camera.tdi_pulses.lte,omitempty"`
	CloudCoverEstimated                   string `url:"cloud_cover.estimated,omitempty"`
	CloudCoverEstimatedGreaterThan        string `url:"cloud_cover.estimated.gt,omitempty"`
	CloudCoverEstimatedGreaterThanOrEqual string `url:"cloud_cover.estimated.gte,omitempty"`
	CloudCoverEstimatedLessThan           string `url:"cloud_cover.estimated.lt,omitempty"`
	CloudCoverEstimatedLessThanOrEqual    string `url:"cloud_cover.estimated.lte,omitempty"`
	FileSize                              string `url:"file_size,omitempty"`
	FileSizeGreaterThan                   string `url:"file_size.gt,omitempty"`
	FileSizeGreaterThanOrEqual            string `url:"file_size.gte,omitempty"`
	FileSizeLessThan                      string `url:"file_size.lt,omitempty"`
	FileSizeLessThanOrEqual               string `url:"file_size.lte,omitempty"`
	ImageStatisticsGSD                    string `url:"image_statistics.gsd,omitempty"`
	ImageStatisticsGSDGreaterThan         string `url:"image_statistics.gsd.gt,omitempty"`
	ImageStatisticsGSDGreaterThanOrEqual  string `url:"image_statistics.gsd.gte,omitempty"`
	ImageStatisticsGSDLessThan            string `url:"image_statistics.gsd.lt,omitempty"`
	ImageStatisticsGSDLessThanOrEqual     string `url:"image_statistics.gsd.lte,omitempty"`
	ImageStatisticsSNR                    string `url:"image_statistics.snr,omitempty"`
	ImageStatisticsSNRGreaterThan         string `url:"image_statistics.snr.gt,omitempty"`
	ImageStatisticsSNRGreaterThanOrEqual  string `url:"image_statistics.snr.gte,omitempty"`
	ImageStatisticsSNRLessThan            string `url:"image_statistics.snr.lt,omitempty"`
	ImageStatisticsSNRLessThanOrEqual     string `url:"image_statistics.snr.lte,omitempty"`
	SatAltitude                           string `url:"sat.alt,omitempty"`
	SatAltitudeGreaterThan                string `url:"sat.alt.gt,omitempty"`
	SatAltitudeGreaterThanOrEqual         string `url:"sat.alt.gte,omitempty"`
	SatAltitudeLessThan                   string `url:"sat.alt.lt,omitempty"`
	SatAltitudeLessThanOrEqual            string `url:"sat.alt.lte,omitempty"`
	SatLat                                string `url:"sat.lat,omitempty"`
	SatLatGreaterThan                     string `url:"sat.lat.gt,omitempty"`
	SatLatGreaterThanOrEqual              string `url:"sat.lat.gte,omitempty"`
	SatLatLessThan                        string `url:"sat.lat.lt,omitempty"`
	SatLatLessThanOrEqual                 string `url:"sat.lat.lte,omitempty"`
	SatLng                                string `url:"sat.lng,omitempty"`
	SatLngGreaterThan                     string `url:"sat.lng.gt,omitempty"`
	SatLngGreaterThanOrEqual              string `url:"sat.lng.gte,omitempty"`
	SatLngLessThan                        string `url:"sat.lng.lt,omitempty"`
	SatLngLessThanOrEqual                 string `url:"sat.lng.lte,omitempty"`
	SatOffNadir                           string `url:"sat.off_nadir,omitempty"`
	SatOffNadirGreaterThan                string `url:"sat.off_nadir.gt,omitempty"`
	SatOffNadirGreaterThanOrEqual         string `url:"sat.off_nadir.gte,omitempty"`
	SatOffNadirLessThan                   string `url:"sat.off_nadir.lt,omitempty"`
	SatOffNadirLessThanOrEqual            string `url:"sat.off_nadir.lte,omitempty"`
	StripId                               string `url:"strip_id,omitempty"`
	StripIdGreaterThan                    string `url:"strip_id.gt,omitempty"`
	StripIdGreaterThanOrEqual             string `url:"strip_id.gte,omitempty"`
	StripIdLessThan                       string `url:"strip_id.lt,omitempty"`
	StripIdLessThanOrEqual                string `url:"strip_id.lte,omitempty"`
	SunAltitude                           string `url:"sun.altitude,omitempty"`
	SunAltitudeGreaterThan                string `url:"sun.altitude.gt,omitempty"`
	SunAltitudeGreaterThanOrEqual         string `url:"sun.altitude.gte,omitempty"`
	SunAltitudeLessThan                   string `url:"sun.altitude.lt,omitempty"`
	SunAltitudeLessThanOrEqual            string `url:"sun.altitude.lte,omitempty"`
	SunAzimuth                            string `url:"sun.azimuth,omitempty"`
	SunAzimuthGreaterThan                 string `url:"sun.azimuth.gt,omitempty"`
	SunAzimuthGreaterThanOrEqual          string `url:"sun.azimuth.gte,omitempty"`
	SunAzimuthLessThan                    string `url:"sun.azimuth.lt,omitempty"`
	SunAzimuthLessThanOrEqual             string `url:"sun.azimuth.lte,omitempty"`
	SunLocalTimeOfDay                     string `url:"sun.local_time_of_day,omitempty"`
	SunLocalTimeOfDayGreaterThan          string `url:"sun.local_time_of_day.gt,omitempty"`
	SunLocalTimeOfDayGreaterThanOrEqual   string `url:"sun.local_time_of_day.gte,omitempty"`
	SunLocalTimeOfDayLessThan             string `url:"sun.local_time_of_day.lt,omitempty"`
	SunLocalTimeOfDayLessThanOrEqual      string `url:"sun.local_time_of_day.lte,omitempty"`
}
