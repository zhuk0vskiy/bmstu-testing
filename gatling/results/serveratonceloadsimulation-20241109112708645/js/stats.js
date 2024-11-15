var stats = {
    type: "GROUP",
name: "Global Information",
path: "",
pathFormatted: "group_missing-name-b06d1",
stats: {
    "name": "Global Information",
    "numberOfRequests": {
        "total": "100000",
        "ok": "99998",
        "ko": "2"
    },
    "minResponseTime": {
        "total": "796",
        "ok": "796",
        "ko": "18230"
    },
    "maxResponseTime": {
        "total": "29133",
        "ok": "29133",
        "ko": "19256"
    },
    "meanResponseTime": {
        "total": "19672",
        "ok": "19672",
        "ko": "18743"
    },
    "standardDeviation": {
        "total": "2727",
        "ok": "2727",
        "ko": "513"
    },
    "percentiles1": {
        "total": "20814",
        "ok": "20815",
        "ko": "18743"
    },
    "percentiles2": {
        "total": "21880",
        "ok": "21882",
        "ko": "19000"
    },
    "percentiles3": {
        "total": "22419",
        "ok": "22419",
        "ko": "19205"
    },
    "percentiles4": {
        "total": "24044",
        "ok": "24043",
        "ko": "19246"
    },
    "group1": {
    "name": "t < 800 ms",
    "count": 20,
    "percentage": 0
},
    "group2": {
    "name": "800 ms < t < 1200 ms",
    "count": 2,
    "percentage": 0
},
    "group3": {
    "name": "t > 1200 ms",
    "count": 99976,
    "percentage": 100
},
    "group4": {
    "name": "failed",
    "count": 2,
    "percentage": 0
},
    "meanNumberOfRequestsPerSecond": {
        "total": "2631.579",
        "ok": "2631.526",
        "ko": "0.053"
    }
},
contents: {
"req_echo-metrics-ed595": {
        type: "REQUEST",
        name: "Echo Metrics",
path: "Echo Metrics",
pathFormatted: "req_echo-metrics-ed595",
stats: {
    "name": "Echo Metrics",
    "numberOfRequests": {
        "total": "50000",
        "ok": "50000",
        "ko": "0"
    },
    "minResponseTime": {
        "total": "796",
        "ok": "796",
        "ko": "-"
    },
    "maxResponseTime": {
        "total": "29133",
        "ok": "29133",
        "ko": "-"
    },
    "meanResponseTime": {
        "total": "19772",
        "ok": "19772",
        "ko": "-"
    },
    "standardDeviation": {
        "total": "2670",
        "ok": "2670",
        "ko": "-"
    },
    "percentiles1": {
        "total": "21133",
        "ok": "21133",
        "ko": "-"
    },
    "percentiles2": {
        "total": "21859",
        "ok": "21859",
        "ko": "-"
    },
    "percentiles3": {
        "total": "22376",
        "ok": "22376",
        "ko": "-"
    },
    "percentiles4": {
        "total": "23843",
        "ok": "23843",
        "ko": "-"
    },
    "group1": {
    "name": "t < 800 ms",
    "count": 19,
    "percentage": 0
},
    "group2": {
    "name": "800 ms < t < 1200 ms",
    "count": 2,
    "percentage": 0
},
    "group3": {
    "name": "t > 1200 ms",
    "count": 49979,
    "percentage": 100
},
    "group4": {
    "name": "failed",
    "count": 0,
    "percentage": 0
},
    "meanNumberOfRequestsPerSecond": {
        "total": "1315.789",
        "ok": "1315.789",
        "ko": "-"
    }
}
    },"req_fasthttp-metric-06f3a": {
        type: "REQUEST",
        name: "FastHTTP Metrics",
path: "FastHTTP Metrics",
pathFormatted: "req_fasthttp-metric-06f3a",
stats: {
    "name": "FastHTTP Metrics",
    "numberOfRequests": {
        "total": "50000",
        "ok": "49998",
        "ko": "2"
    },
    "minResponseTime": {
        "total": "797",
        "ok": "797",
        "ko": "18230"
    },
    "maxResponseTime": {
        "total": "29133",
        "ok": "29133",
        "ko": "19256"
    },
    "meanResponseTime": {
        "total": "19572",
        "ok": "19572",
        "ko": "18743"
    },
    "standardDeviation": {
        "total": "2779",
        "ok": "2779",
        "ko": "513"
    },
    "percentiles1": {
        "total": "20250",
        "ok": "20250",
        "ko": "18743"
    },
    "percentiles2": {
        "total": "21901",
        "ok": "21901",
        "ko": "19000"
    },
    "percentiles3": {
        "total": "22486",
        "ok": "22486",
        "ko": "19205"
    },
    "percentiles4": {
        "total": "24236",
        "ok": "24236",
        "ko": "19246"
    },
    "group1": {
    "name": "t < 800 ms",
    "count": 1,
    "percentage": 0
},
    "group2": {
    "name": "800 ms < t < 1200 ms",
    "count": 0,
    "percentage": 0
},
    "group3": {
    "name": "t > 1200 ms",
    "count": 49997,
    "percentage": 100
},
    "group4": {
    "name": "failed",
    "count": 2,
    "percentage": 0
},
    "meanNumberOfRequestsPerSecond": {
        "total": "1315.789",
        "ok": "1315.737",
        "ko": "0.053"
    }
}
    }
}

}

function fillStats(stat){
    $("#numberOfRequests").append(stat.numberOfRequests.total);
    $("#numberOfRequestsOK").append(stat.numberOfRequests.ok);
    $("#numberOfRequestsKO").append(stat.numberOfRequests.ko);

    $("#minResponseTime").append(stat.minResponseTime.total);
    $("#minResponseTimeOK").append(stat.minResponseTime.ok);
    $("#minResponseTimeKO").append(stat.minResponseTime.ko);

    $("#maxResponseTime").append(stat.maxResponseTime.total);
    $("#maxResponseTimeOK").append(stat.maxResponseTime.ok);
    $("#maxResponseTimeKO").append(stat.maxResponseTime.ko);

    $("#meanResponseTime").append(stat.meanResponseTime.total);
    $("#meanResponseTimeOK").append(stat.meanResponseTime.ok);
    $("#meanResponseTimeKO").append(stat.meanResponseTime.ko);

    $("#standardDeviation").append(stat.standardDeviation.total);
    $("#standardDeviationOK").append(stat.standardDeviation.ok);
    $("#standardDeviationKO").append(stat.standardDeviation.ko);

    $("#percentiles1").append(stat.percentiles1.total);
    $("#percentiles1OK").append(stat.percentiles1.ok);
    $("#percentiles1KO").append(stat.percentiles1.ko);

    $("#percentiles2").append(stat.percentiles2.total);
    $("#percentiles2OK").append(stat.percentiles2.ok);
    $("#percentiles2KO").append(stat.percentiles2.ko);

    $("#percentiles3").append(stat.percentiles3.total);
    $("#percentiles3OK").append(stat.percentiles3.ok);
    $("#percentiles3KO").append(stat.percentiles3.ko);

    $("#percentiles4").append(stat.percentiles4.total);
    $("#percentiles4OK").append(stat.percentiles4.ok);
    $("#percentiles4KO").append(stat.percentiles4.ko);

    $("#meanNumberOfRequestsPerSecond").append(stat.meanNumberOfRequestsPerSecond.total);
    $("#meanNumberOfRequestsPerSecondOK").append(stat.meanNumberOfRequestsPerSecond.ok);
    $("#meanNumberOfRequestsPerSecondKO").append(stat.meanNumberOfRequestsPerSecond.ko);
}
