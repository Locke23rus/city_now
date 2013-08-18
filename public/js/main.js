var feed = new EventSource('/feed');

var updateStats = function(msg) {
  var stats = JSON.parse(msg.data);
  $('#online').text(stats.online);
}

feed.addEventListener("stats", updateStats, false);
