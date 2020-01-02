var settings = {
	"async": true,
	"crossDomain": true,
	"url": "https://cometari-airportsfinder-v1.p.rapidapi.com/api/airports/by-text?text=Bangkok",
	"method": "GET",
	"headers": {
		"x-rapidapi-host": "cometari-airportsfinder-v1.p.rapidapi.com",
		"x-rapidapi-key": "93be61accemsh1c6e7979e10e3c7p1ad963jsn20338522fb9d"
	}
}

$.ajax(settings).done(function (response) {
	console.log(response);
});