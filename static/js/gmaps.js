
function  getGeoPos() {
  // `https://maps.googleapis.com/maps/api/geocode/json?address=1600+Amphitheatre+Parkway,+Mountain+View,+CA&key=AIzaSyDic50ciJOPIZ05TViE1E7NyiHEtkVWcxo`
}

function initMap() {

  var myLatLng = {lat: 46.00, lng: 2.00};
  // Create a map object and specify the DOM element for display.
  var map = new google.maps.Map(document.getElementById('map'), {
    center: myLatLng,
    scrollwheel: false,
    zoom: 4
  });

  var request = {
    location: map.getCenter(),
    radius: '500',
    query: $('#adress').html()
  };

  var service = new google.maps.places.PlacesService(map);
  service.textSearch(request, maprefresh);
}

function maprefresh(results, status) {
  if (status == google.maps.places.PlacesServiceStatus.OK) {
    if(results.length != 0){
      var place = results[0];
      console.log(place.geometry.location.lat(), place.geometry.location.lng())
      var myLatLng = {lat: place.geometry.location.lat(), lng: place.geometry.location.lng()};
      map = new google.maps.Map(document.getElementById('map'), {
        center: myLatLng,
        scrollwheel: false,
        zoom: 14
      });
      var marker = new google.maps.Marker({
        map: map,
        position: myLatLng
      });
    }
    else{
      var myLatLng = {lat: 46.00, lng: 2.00};
      map = new google.maps.Map(document.getElementById('map'), {
        center: myLatLng,
        scrollwheel: false,
        zoom: 14
      });
    }
  }
}
