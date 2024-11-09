var stats = {
    type: "GROUP",
name: "Global Information",
path: "",
pathFormatted: "group_missing-name-b06d1",
stats: {
    "name": "Global Information",
    "numberOfRequests": {
        "total": "100000",
        "ok": "100000",
        "ko": "0"
    },
    "minResponseTime": {
        "total": "704",
        "ok": "704",
        "ko": "-"
    },
    "maxResponseTime": {
        "total": "30572",
        "ok": "30572",
        "ko": "-"
    },
    "meanResponseTime": {
        "total": "22208",
        "ok": "22208",
        "ko": "-"
    },
    "standardDeviation": {
        "total": "2901",
        "ok": "2901",
        "ko": "-"
    },
    "percentiles1": {
        "total": "22945",
        "ok": "22946",
        "ko": "-"
    },
    "percentiles2": {
        "total": "24691",
        "ok": "24691",
        "ko": "-"
    },
    "percentiles3": {
        "total": "25380",
        "ok": "25380",
        "ko": "-"
    },
    "percentiles4": {
        "total": "26147",
        "ok": "26147",
        "ko": "-"
    },
    "group1": {
    "name": "t < 800 ms",
    "count": 22,
    "percentage": 0
},
    "group2": {
    "name": "800 ms < t < 1200 ms",
    "count": 0,
    "percentage": 0
},
    "group3": {
    "name": "t > 1200 ms",
    "count": 99978,
    "percentage": 100
},
    "group4": {
    "name": "failed",
    "count": 0,
    "percentage": 0
},
    "meanNumberOfRequestsPerSecond": {
        "total": "2439.024",
        "ok": "2439.024",
        "ko": "-"
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
        "total": "704",
        "ok": "704",
        "ko": "-"
    },
    "maxResponseTime": {
        "total": "30551",
        "ok": "30551",
        "ko": "-"
    },
    "meanResponseTime": {
        "total": "22155",
        "ok": "22155",
        "ko": "-"
    },
    "standardDeviation": {
        "total": "2926",
        "ok": "2926",
        "ko": "-"
    },
    "percentiles1": {
        "total": "22867",
        "ok": "22867",
        "ko": "-"
    },
    "percentiles2": {
        "total": "24664",
        "ok": "24664",
        "ko": "-"
    },
    "percentiles3": {
        "total": "25288",
        "ok": "25288",
        "ko": "-"
    },
    "percentiles4": {
        "total": "26226",
        "ok": "26226",
        "ko": "-"
    },
    "group1": {
    "name": "t < 800 ms",
    "count": 21,
    "percentage": 0
},
    "group2": {
    "name": "800 ms < t < 1200 ms",
    "count": 0,
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
        "total": "1219.512",
        "ok": "1219.512",
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
        "ok": "50000",
        "ko": "0"
    },
    "minResponseTime": {
        "total": "718",
        "ok": "718",
        "ko": "-"
    },
    "maxResponseTime": {
        "total": "30572",
        "ok": "30572",
        "ko": "-"
    },
    "meanResponseTime": {
        "total": "22262",
        "ok": "22262",
        "ko": "-"
    },
    "standardDeviation": {
        "total": "2876",
        "ok": "2876",
        "ko": "-"
    },
    "percentiles1": {
        "total": "23009",
        "ok": "23009",
        "ko": "-"
    },
    "percentiles2": {
        "total": "24729",
        "ok": "24729",
        "ko": "-"
    },
    "percentiles3": {
        "total": "25408",
        "ok": "25408",
        "ko": "-"
    },
    "percentiles4": {
        "total": "26128",
        "ok": "26128",
        "ko": "-"
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
    "count": 49999,
    "percentage": 100
},
    "group4": {
    "name": "failed",
    "count": 0,
    "percentage": 0
},
    "meanNumberOfRequestsPerSecond": {
        "total": "1219.512",
        "ok": "1219.512",
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
