package httpclient

import "time"


type HttpClient struct {}

func NewHttpClient() *HttpClient {
	return &HttpClient{}
}

func (h *HttpClient) GetWeatherData(longitude float64, latitude float64) (string, error) {
	time.Sleep(time.Millisecond * 25)

	if longitude > -20 && longitude < 50 && latitude > 50 {

		// Burası Kuzey Avrupa'ya olabilir.

		return `Yakın geleceğinde yağmur ve rüzgar var, dostum. Bugün muhtemelen yağmur yağacak. 
		Yarın da muhtemelen yağmur yağacak. Ayrıca, bahar denen şeyi hiç duydun mu? 
		Senin için değil! Kış geliyor, ebedi kış. Güç, dostum.`, nil
	}
	return `Güneş şu anda ışıldamıyor olsa bile, yakında ışıma ihtimali çok yüksek! İşte orada 
	sizin ve sevdikleriniz için umuttur, belki bir noktada sahile bile gidebilirsiniz 
	ve okyanus dalgalarının kıyıya vuran yumuşak seslerini dinlerken kokteyllerinizi yudumlayın. 
	Dünya çok güzel bir yer...`, nil
}