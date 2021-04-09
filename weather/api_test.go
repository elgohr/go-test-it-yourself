package weather_test

import (
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"workshop/weather"
)

func TestCurrentTemperature(t *testing.T) {
	// see https://weatherstack.com/documentation#current_weather
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/current", r.URL.Path)
		require.Equal(t, "Wolfsburg", r.URL.Query().Get("query"))

		testContent, err := ioutil.ReadFile("testdata/current.json")
		require.NoError(t, err)
		w.Write(testContent)
	}))
	defer ts.Close()

	temperature, err := weather.CurrentTemperature(ts.URL, "Wolfsburg")
	require.NoError(t, err)
	require.Equal(t, 16, temperature)
}

func TestCurrentTemperatureFailingRequest(t *testing.T) {
	temperature, err := weather.CurrentTemperature("http://localhost/nothing_here", "Wolfsburg")
	require.EqualError(t, err, "could not contact api")
	require.Equal(t, -90, temperature)
}

func TestCurrentTemperatureFailingDecode(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("I'm serving everything but the weather"))
	}))
	defer ts.Close()

	temperature, err := weather.CurrentTemperature(ts.URL, "Wolfsburg")
	require.EqualError(t, err, "could not decode weather")
	require.Equal(t, -90, temperature)
}
