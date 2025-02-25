var map = L.map('map').setView(Toulouse, 1.5);

L.tileLayer('https://tile.openstreetmap.org/{z}/{x}/{y}.png', {
    maxZoom: 19,
    attribution: '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>'
}).addTo(map);

for (let x=0; x < AllCoordinates.length; x++) {
    L.marker(AllCoordinates[x]).addTo(map)
        .bindPopup(AllLocations[x]);
}