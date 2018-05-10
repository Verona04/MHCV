// Globale variabler vi benytter overalt:
var features = {}
var openFeature = null
var controls = {}
var zoom = 12
var moveTimeout = null
var vectorLayer = null
var enableRecenter = true

var map = new OpenLayers.Map({
    div: "mapdiv",
    eventListeners: {
        featureclick: function (e) {
            var parkid = e.feature.data.parkid
            if (parkid) {
                $('.line-parkid-' + parkid)[0].scrollIntoView()
                $('.line-parkid-' + parkid).effect("highlight", {}, 1000)
            }
        }
    }
})
map.events.register('click', map, function (e) {
    var position = map.getLonLatFromPixel(e.xy);
    var lonLat = new OpenLayers.LonLat(position.lon, position.lat)
        .transform(map.getProjectionObject(), new OpenLayers.Projection("EPSG:4326"))

    getParking(1500, lonLat.lon, lonLat.lat)
})

// jquery-funksjon som kjører i det vi åpner siden
$(function() {
    map.addLayer(new OpenLayers.Layer.OSM())
    addMoveEventListener()
    myPosition(null)
    initOpenLayersMap()
})
// Funksjon som gjør parkeringsplass-søk når vi flytter rundt på kartet.
function addMoveEventListener () {
    map.events.listeners.move.unshift({func: function(ev) {

           /* clearTimeout(moveTimeout)
            moveTimeout = setTimeout(function() {
                var position = map.getCenter()
                var lonLat = new OpenLayers.LonLat(position.lon, position.lat)
                    .transform(map.getProjectionObject(), new OpenLayers.Projection("EPSG:4326"))
                var parkRadius = 1500 //document.getElementById('parkRadius').value
                getParking(parkRadius, lonLat.lon, lonLat.lat)
            }, 500)
            */

        }})
}

function areasearch(event) {
    enableRecenter = false
    var position = map.getCenter()
    var lonLat = new OpenLayers.LonLat(position.lon, position.lat)
        .transform(map.getProjectionObject(), new OpenLayers.Projection("EPSG:4326"))
    getParking(10000, lonLat.lon, lonLat.lat)


    $('#searchResult').animate({scrollTop: 0})

    event.preventDefault()
    event.stopPropagation()
}

// Funksjon for å sentrere kartet til nettleserens posisjon.
function myPosition (event) {
    if ("geolocation" in navigator) {
        navigator.geolocation.getCurrentPosition(function(position) {
            var fromProjection = new OpenLayers.Projection("EPSG:4326") // transform from WGS 1984
            var toProjection = map.getProjectionObject() // to Spherical Mercator Projection
            var lonLat = new OpenLayers.LonLat(position.coords.longitude, position.coords.latitude)
                .transform(fromProjection, toProjection)
            map.setCenter(lonLat, zoom)
        })
    }
    if (event !== null) {
        event.stopPropagation()
        event.preventDefault()
    }
    return false
}
// Behandling av søkeresultat.
function parkingSearchResult (res) {
    if (res == null) {
        return
    }
    var ft = Object.keys(features)
    if (ft !== null) {
        ft.map(function (k) {
            vectorLayer.removeFeatures(features[k])
        })
    }
    res.map(function(park) {
        features[park.id] = new OpenLayers.Feature.Vector(
            new OpenLayers.Geometry.Point(park.lengdegrad,park.breddegrad)
                .transform(new OpenLayers.Projection("EPSG:4326"), map.getProjectionObject()),
            {description:park.aktivVersjon.navn, parkid: park.id}
        )
        vectorLayer.addFeatures(features[park.id])
    })

    lagResultatListe(res)
}
// Oppslag for parkering innenfor et gitt område oppgitt i lengde og breddegrader.
function getParking(radius, longitude, latitude) {
    var hc = document.getElementById('hc').checked ? 'on' : ''
    var ladestasjoner = document.getElementById('ladestasjoner').checked ? 'on' : ''
    if (!radius) {
        radius = 1500
    }
    $.get(
        '/api/parkering/radius?radius=' + radius +
        '&longitude=' + longitude +
        '&latitude=' + latitude+'&hc=' + hc + '&ladestasjoner=' + ladestasjoner).then(parkingSearchResult)
}
// Oppslag for parkering innenfor en gitt søketekst.
function searchParking (event) {
    enableRecenter = true
    var search = document.getElementById('search').value
    $.get('/api/parkering/search?search=' + search).then(parkingSearchResult)
    event.preventDefault()
    event.stopPropagation()
}
// Highlighter en spesifik parkeringsplass
function centerOnMap (id, breddegrad, lengdegrad, zoom) {
    var lonLat = new OpenLayers.LonLat(lengdegrad, breddegrad)
    lonLat.transform(
        new OpenLayers.Projection("EPSG:4326"), // transform from WGS 1984
        map.getProjectionObject() // to Spherical Mercator Projection
    )
    map.setCenter(lonLat, zoom)
    if (zoom == 16) {
        //createPopup(features[id])
    }
}
// Lager en popup i openlayers.
function createPopup (feature) {
    if (openFeature !== null) {
        destroyPopup(openFeature)
    }
    openFeature = feature
    feature.popup = new OpenLayers.Popup.FramedCloud(
        "pop",
        feature.geometry.getBounds().getCenterLonLat(),
        null,
        '<div class="markerContent">'+feature.attributes.description+'</div>',
        null,
        true,
        function() { controls['selector'].unselectAll(); }
    )
    //feature.popup.closeOnMove = true
    map.addPopup(feature.popup)
}
// Sletter en popup i openlayers
function destroyPopup (feature) {
    feature.popup.destroy()
    feature.popup = null
    openFeature = null
}
function initOpenLayersMap () {
    vectorLayer = new OpenLayers.Layer.Vector("Overlay")
    map.addLayer(vectorLayer)
    controls = {
        selector: new OpenLayers.Control.SelectFeature(
            vectorLayer,
            {
                onSelect: createPopup,
                onUnselect: destroyPopup
            }
        )
    }
    map.addControl(controls['selector'])
    controls['selector'].activate()
}

function runSearch (event) {
    var searchTerm = document.getElementById('search').value
    var hc = document.getElementById('hc').checked ? 'on' : ''
    var ladestasjoner = document.getElementById('ladestasjoner').checked ? 'on' : ''

    $.get(
        '/api/parkering/search?search=' + searchTerm + '&hc=' + hc + '&ladestasjoner=' + ladestasjoner
    ).then(parkingSearchResult)

    event.stopPropagation()
    event.preventDefault()
    return false
}

function lagResultatListe (resultater) {

    if (enableRecenter) {
        centerOnMap(resultater[0].id, resultater[0].breddegrad, resultater[0].lengdegrad, 11)
    }

    var searchResults = $('#searchResult')
    searchResults.empty()
    resultater.map(function(park) {

        var line = $(`
<div class="row line-parkid-${park.id}" style="border: 2px solid #000; margin: 10px;">
    <div class="col-md-12">
        <div class="row">
            <div class="col-md-10 parkering-navn">
                <b>${park.aktivVersjon.navn}</b>
            </div>
            <div class="col-md-2" style="text-align: right;">
                <a href="#" onClick="centerOnMap(${park.id}, ${park.breddegrad}, ${park.lengdegrad}, 16)">Vis på kart</a>
            </div>
        </div>

        <div class="row parkring-type">
            <div class="col-md-4">
                ${park.aktivVersjon.typeParkeringsomrade}<br>
            </div>
            <div class="col-md-8">
                ${park.aktivVersjon.adresse} ${park.aktivVersjon.postnummer} ${park.aktivVersjon.poststed}<br>
            </div>
        </div>

        <div class="row">
            <div class="col-md-4">
                ${park.aktivVersjon.antallAvgiftsfriePlasser} Antall Avgiftsfrie Plasser<br>
                ${park.aktivVersjon.antallAvgiftsbelagtePlasser} Antall Avgiftsbelagte Plasser<br>
                ${park.aktivVersjon.antallLadeplasser} Antall Ladeplasser<br>
            </div>

            <div class="col-md-8">
                                   
                ${park.aktivVersjon.antallForflytningshemmede} Antall Forflytnigshemmede<br>
                ${park.aktivVersjon.vurderingForflytningshemmede}<br>
            </div>
        </div>
    </div>
</div>
`)

        searchResults.append(line)
    })
}
