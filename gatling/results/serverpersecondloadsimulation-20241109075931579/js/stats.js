var stats = {
    type: "GROUP",
name: "Global Information",
path: "",
pathFormatted: "group_missing-name-b06d1",
stats: {
    "name": "Global Information",
    "numberOfRequests": {
        "total": "120000",
        "ok": "119792",
        "ko": "208"
    },
    "minResponseTime": {
        "total": "848",
        "ok": "848",
        "ko": "12577"
    },
    "maxResponseTime": {
        "total": "33874",
        "ok": "33874",
        "ko": "27335"
    },
    "meanResponseTime": {
        "total": "24399",
        "ok": "24401",
        "ko": "23310"
    },
    "standardDeviation": {
        "total": "4026",
        "ok": "4028",
        "ko": "3008"
    },
    "percentiles1": {
        "total": "24274",
        "ok": "24274",
        "ko": "23891"
    },
    "percentiles2": {
        "total": "27864",
        "ok": "27866",
        "ko": "24528"
    },
    "percentiles3": {
        "total": "30281",
        "ok": "30286",
        "ko": "26861"
    },
    "percentiles4": {
        "total": "31461",
        "ok": "31469",
        "ko": "26987"
    },
    "group1": {
    "name": "t < 800 ms",
    "count": 0,
    "percentage": 0
},
    "group2": {
    "name": "800 ms < t < 1200 ms",
    "count": 18,
    "percentage": 0
},
    "group3": {
    "name": "t > 1200 ms",
    "count": 119774,
    "percentage": 100
},
    "group4": {
    "name": "failed",
    "count": 208,
    "percentage": 0
},
    "meanNumberOfRequestsPerSecond": {
        "total": "2500",
        "ok": "2495.667",
        "ko": "4.333"
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
        "ok": "59792",
        "ko": "208"
    },
    "minResponseTime": {
        "total": "848",
        "ok": "848",
        "ko": "12577"
    },
    "maxResponseTime": {
        "total": "33874",
        "ok": "33874",
        "ko": "27335"
    },
    "meanResponseTime": {
        "total": "24329",
        "ok": "24333",
        "ko": "23310"
    },
    "standardDeviation": {
        "total": "4006",
        "ok": "4008",
        "ko": "3008"
    },
    "percentiles1": {
        "total": "24202",
        "ok": "24206",
        "ko": "23891"
    },
    "percentiles2": {
        "total": "27605",
        "ok": "27620",
        "ko": "24528"
    },
    "percentiles3": {
        "total": "30231",
        "ok": "30233",
        "ko": "26861"
    },
    "percentiles4": {
        "total": "31426",
        "ok": "31425",
        "ko": "26987"
    },
    "group1": {
    "name": "t < 800 ms",
    "count": 0,
    "percentage": 0
},
    "group2": {
    "name": "800 ms < t < 1200 ms",
    "count": 18,
    "percentage": 0
},
    "group3": {
    "name": "t > 1200 ms",
    "count": 59774,
    "percentage": 100
},
    "group4": {
    "name": "failed",
    "count": 208,
    "percentage": 0
},
    "meanNumberOfRequestsPerSecond": {
        "total": "1250",
        "ok": "1245.667",
        "ko": "4.333"
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
        "ok": "60000",
        "ko": "0"
    },
    "minResponseTime": {
        "total": "6026",
        "ok": "6026",
        "ko": "-"
    },
    "maxResponseTime": {
        "total": "33874",
        "ok": "33874",
        "ko": "-"
    },
    "meanResponseTime": {
        "total": "24470",
        "ok": "24470",
        "ko": "-"
    },
    "standardDeviation": {
        "total": "4046",
        "ok": "4046",
        "ko": "-"
    },
    "percentiles1": {
        "total": "24295",
        "ok": "24296",
        "ko": "-"
    },
    "percentiles2": {
        "total": "27987",
        "ok": "27986",
        "ko": "-"
    },
    "percentiles3": {
        "total": "30338",
        "ok": "30338",
        "ko": "-"
    },
    "percentiles4": {
        "total": "31519",
        "ok": "31529",
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
    "count": 60000,
    "percentage": 100
},
    "group4": {
    "name": "failed",
    "count": 0,
    "percentage": 0
},
    "meanNumberOfRequestsPerSecond": {
        "total": "1250",
        "ok": "1250",
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
