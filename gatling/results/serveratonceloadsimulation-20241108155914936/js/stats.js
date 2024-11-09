var stats = {
    type: "GROUP",
name: "Global Information",
path: "",
pathFormatted: "group_missing-name-b06d1",
stats: {
    "name": "Global Information",
    "numberOfRequests": {
        "total": "100000",
        "ok": "99984",
        "ko": "16"
    },
    "minResponseTime": {
        "total": "926",
        "ok": "926",
        "ko": "18773"
    },
    "maxResponseTime": {
        "total": "30650",
        "ok": "30650",
        "ko": "20533"
    },
    "meanResponseTime": {
        "total": "24070",
        "ok": "24071",
        "ko": "19437"
    },
    "standardDeviation": {
        "total": "4261",
        "ok": "4261",
        "ko": "473"
    },
    "percentiles1": {
        "total": "26167",
        "ok": "26169",
        "ko": "19382"
    },
    "percentiles2": {
        "total": "27791",
        "ok": "27790",
        "ko": "19621"
    },
    "percentiles3": {
        "total": "28431",
        "ok": "28431",
        "ko": "20443"
    },
    "percentiles4": {
        "total": "28800",
        "ok": "28802",
        "ko": "20515"
    },
    "group1": {
    "name": "t < 800 ms",
    "count": 0,
    "percentage": 0
},
    "group2": {
    "name": "800 ms < t < 1200 ms",
    "count": 22,
    "percentage": 0
},
    "group3": {
    "name": "t > 1200 ms",
    "count": 99962,
    "percentage": 100
},
    "group4": {
    "name": "failed",
    "count": 16,
    "percentage": 0
},
    "meanNumberOfRequestsPerSecond": {
        "total": "2380.952",
        "ok": "2380.571",
        "ko": "0.381"
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
        "ok": "49984",
        "ko": "16"
    },
    "minResponseTime": {
        "total": "926",
        "ok": "926",
        "ko": "18773"
    },
    "maxResponseTime": {
        "total": "30650",
        "ok": "30650",
        "ko": "20533"
    },
    "meanResponseTime": {
        "total": "23671",
        "ok": "23673",
        "ko": "19437"
    },
    "standardDeviation": {
        "total": "4282",
        "ok": "4282",
        "ko": "473"
    },
    "percentiles1": {
        "total": "25477",
        "ok": "25480",
        "ko": "19382"
    },
    "percentiles2": {
        "total": "27611",
        "ok": "27611",
        "ko": "19621"
    },
    "percentiles3": {
        "total": "28428",
        "ok": "28428",
        "ko": "20443"
    },
    "percentiles4": {
        "total": "28818",
        "ok": "28818",
        "ko": "20515"
    },
    "group1": {
    "name": "t < 800 ms",
    "count": 0,
    "percentage": 0
},
    "group2": {
    "name": "800 ms < t < 1200 ms",
    "count": 22,
    "percentage": 0
},
    "group3": {
    "name": "t > 1200 ms",
    "count": 49962,
    "percentage": 100
},
    "group4": {
    "name": "failed",
    "count": 16,
    "percentage": 0
},
    "meanNumberOfRequestsPerSecond": {
        "total": "1190.476",
        "ok": "1190.095",
        "ko": "0.381"
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
        "ok": "50000",
        "ko": "0"
    },
    "minResponseTime": {
        "total": "5098",
        "ok": "5098",
        "ko": "-"
    },
    "maxResponseTime": {
        "total": "30420",
        "ok": "30420",
        "ko": "-"
    },
    "meanResponseTime": {
        "total": "24470",
        "ok": "24470",
        "ko": "-"
    },
    "standardDeviation": {
        "total": "4203",
        "ok": "4203",
        "ko": "-"
    },
    "percentiles1": {
        "total": "26809",
        "ok": "26809",
        "ko": "-"
    },
    "percentiles2": {
        "total": "27931",
        "ok": "27931",
        "ko": "-"
    },
    "percentiles3": {
        "total": "28432",
        "ok": "28432",
        "ko": "-"
    },
    "percentiles4": {
        "total": "28779",
        "ok": "28779",
        "ko": "-"
    },
    "group1": {
    "name": "t < 800 ms",
    "count": 0,
    "percentage": 0
},
    "group2": {
    "name": "800 ms < t < 1200 ms",
    "count": 0,
    "percentage": 0
},
    "group3": {
    "name": "t > 1200 ms",
    "count": 50000,
    "percentage": 100
},
    "group4": {
    "name": "failed",
    "count": 0,
    "percentage": 0
},
    "meanNumberOfRequestsPerSecond": {
        "total": "1190.476",
        "ok": "1190.476",
        "ko": "-"
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
