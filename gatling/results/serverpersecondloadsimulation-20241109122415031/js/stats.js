var stats = {
    type: "GROUP",
name: "Global Information",
path: "",
pathFormatted: "group_missing-name-b06d1",
stats: {
    "name": "Global Information",
    "numberOfRequests": {
        "total": "120000",
        "ok": "0",
        "ko": "120000"
    },
    "minResponseTime": {
        "total": "318",
        "ok": "-",
        "ko": "318"
    },
    "maxResponseTime": {
        "total": "26771",
        "ok": "-",
        "ko": "26771"
    },
    "meanResponseTime": {
        "total": "15553",
        "ok": "-",
        "ko": "15553"
    },
    "standardDeviation": {
        "total": "3936",
        "ok": "-",
        "ko": "3936"
    },
    "percentiles1": {
        "total": "15417",
        "ok": "-",
        "ko": "15417"
    },
    "percentiles2": {
        "total": "17966",
        "ok": "-",
        "ko": "17966"
    },
    "percentiles3": {
        "total": "23167",
        "ok": "-",
        "ko": "23164"
    },
    "percentiles4": {
        "total": "25360",
        "ok": "-",
        "ko": "25360"
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
    "count": 0,
    "percentage": 0
},
    "group4": {
    "name": "failed",
    "count": 120000,
    "percentage": 100
},
    "meanNumberOfRequestsPerSecond": {
        "total": "3000",
        "ok": "-",
        "ko": "3000"
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
        "total": "60000",
        "ok": "0",
        "ko": "60000"
    },
    "minResponseTime": {
        "total": "318",
        "ok": "-",
        "ko": "318"
    },
    "maxResponseTime": {
        "total": "26766",
        "ok": "-",
        "ko": "26766"
    },
    "meanResponseTime": {
        "total": "15583",
        "ok": "-",
        "ko": "15583"
    },
    "standardDeviation": {
        "total": "4011",
        "ok": "-",
        "ko": "4011"
    },
    "percentiles1": {
        "total": "15422",
        "ok": "-",
        "ko": "15423"
    },
    "percentiles2": {
        "total": "17950",
        "ok": "-",
        "ko": "17950"
    },
    "percentiles3": {
        "total": "23258",
        "ok": "-",
        "ko": "23258"
    },
    "percentiles4": {
        "total": "25823",
        "ok": "-",
        "ko": "25823"
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
    "count": 0,
    "percentage": 0
},
    "group4": {
    "name": "failed",
    "count": 60000,
    "percentage": 100
},
    "meanNumberOfRequestsPerSecond": {
        "total": "1500",
        "ok": "-",
        "ko": "1500"
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
        "total": "60000",
        "ok": "0",
        "ko": "60000"
    },
    "minResponseTime": {
        "total": "1826",
        "ok": "-",
        "ko": "1826"
    },
    "maxResponseTime": {
        "total": "26771",
        "ok": "-",
        "ko": "26771"
    },
    "meanResponseTime": {
        "total": "15523",
        "ok": "-",
        "ko": "15523"
    },
    "standardDeviation": {
        "total": "3858",
        "ok": "-",
        "ko": "3858"
    },
    "percentiles1": {
        "total": "15412",
        "ok": "-",
        "ko": "15412"
    },
    "percentiles2": {
        "total": "17978",
        "ok": "-",
        "ko": "17978"
    },
    "percentiles3": {
        "total": "23065",
        "ok": "-",
        "ko": "23065"
    },
    "percentiles4": {
        "total": "25107",
        "ok": "-",
        "ko": "25107"
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
    "count": 0,
    "percentage": 0
},
    "group4": {
    "name": "failed",
    "count": 60000,
    "percentage": 100
},
    "meanNumberOfRequestsPerSecond": {
        "total": "1500",
        "ok": "-",
        "ko": "1500"
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
