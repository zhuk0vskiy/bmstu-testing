var stats = {
    type: "GROUP",
name: "Global Information",
path: "",
pathFormatted: "group_missing-name-b06d1",
stats: {
    "name": "Global Information",
    "numberOfRequests": {
        "total": "113230",
        "ok": "112203",
        "ko": "1027"
    },
    "minResponseTime": {
        "total": "509",
        "ok": "509",
        "ko": "15015"
    },
    "maxResponseTime": {
        "total": "39615",
        "ok": "39615",
        "ko": "38496"
    },
    "meanResponseTime": {
        "total": "28768",
        "ok": "28795",
        "ko": "25839"
    },
    "standardDeviation": {
        "total": "4719",
        "ok": "4705",
        "ko": "5317"
    },
    "percentiles1": {
        "total": "29277",
        "ok": "29307",
        "ko": "25982"
    },
    "percentiles2": {
        "total": "32540",
        "ok": "32557",
        "ko": "30066"
    },
    "percentiles3": {
        "total": "35200",
        "ok": "35208",
        "ko": "34738"
    },
    "percentiles4": {
        "total": "38493",
        "ok": "38499",
        "ko": "38452"
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
    "count": 112181,
    "percentage": 99
},
    "group4": {
    "name": "failed",
    "count": 1027,
    "percentage": 1
},
    "meanNumberOfRequestsPerSecond": {
        "total": "1856.23",
        "ok": "1839.393",
        "ko": "16.836"
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
        "total": "56708",
        "ok": "55740",
        "ko": "968"
    },
    "minResponseTime": {
        "total": "509",
        "ok": "509",
        "ko": "15015"
    },
    "maxResponseTime": {
        "total": "39607",
        "ok": "39607",
        "ko": "38496"
    },
    "meanResponseTime": {
        "total": "28404",
        "ok": "28449",
        "ko": "25836"
    },
    "standardDeviation": {
        "total": "4741",
        "ok": "4716",
        "ko": "5384"
    },
    "percentiles1": {
        "total": "28867",
        "ok": "28911",
        "ko": "25963"
    },
    "percentiles2": {
        "total": "32268",
        "ok": "32290",
        "ko": "30179"
    },
    "percentiles3": {
        "total": "35022",
        "ok": "35044",
        "ko": "34762"
    },
    "percentiles4": {
        "total": "38494",
        "ok": "38505",
        "ko": "38457"
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
    "count": 55718,
    "percentage": 98
},
    "group4": {
    "name": "failed",
    "count": 968,
    "percentage": 2
},
    "meanNumberOfRequestsPerSecond": {
        "total": "929.639",
        "ok": "913.77",
        "ko": "15.869"
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
        "total": "56522",
        "ok": "56463",
        "ko": "59"
    },
    "minResponseTime": {
        "total": "14425",
        "ok": "14425",
        "ko": "17683"
    },
    "maxResponseTime": {
        "total": "39615",
        "ok": "39615",
        "ko": "31033"
    },
    "meanResponseTime": {
        "total": "29133",
        "ok": "29136",
        "ko": "25898"
    },
    "standardDeviation": {
        "total": "4669",
        "ok": "4669",
        "ko": "4067"
    },
    "percentiles1": {
        "total": "30100",
        "ok": "30110",
        "ko": "27827"
    },
    "percentiles2": {
        "total": "32913",
        "ok": "32915",
        "ko": "29404"
    },
    "percentiles3": {
        "total": "35237",
        "ok": "35237",
        "ko": "30029"
    },
    "percentiles4": {
        "total": "38490",
        "ok": "38492",
        "ko": "30715"
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
    "count": 56463,
    "percentage": 100
},
    "group4": {
    "name": "failed",
    "count": 59,
    "percentage": 0
},
    "meanNumberOfRequestsPerSecond": {
        "total": "926.59",
        "ok": "925.623",
        "ko": "0.967"
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
