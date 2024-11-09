var stats = {
    type: "GROUP",
name: "Global Information",
path: "",
pathFormatted: "group_missing-name-b06d1",
stats: {
    "name": "Global Information",
    "numberOfRequests": {
        "total": "120000",
        "ok": "119331",
        "ko": "669"
    },
    "minResponseTime": {
        "total": "454",
        "ok": "454",
        "ko": "15238"
    },
    "maxResponseTime": {
        "total": "33127",
        "ok": "33127",
        "ko": "27651"
    },
    "meanResponseTime": {
        "total": "25007",
        "ok": "25019",
        "ko": "22815"
    },
    "standardDeviation": {
        "total": "3558",
        "ok": "3558",
        "ko": "2815"
    },
    "percentiles1": {
        "total": "25086",
        "ok": "25101",
        "ko": "23606"
    },
    "percentiles2": {
        "total": "27940",
        "ok": "27951",
        "ko": "24722"
    },
    "percentiles3": {
        "total": "29947",
        "ok": "29957",
        "ko": "27357"
    },
    "percentiles4": {
        "total": "31750",
        "ok": "31752",
        "ko": "27525"
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
    "count": 119309,
    "percentage": 99
},
    "group4": {
    "name": "failed",
    "count": 669,
    "percentage": 1
},
    "meanNumberOfRequestsPerSecond": {
        "total": "2448.98",
        "ok": "2435.327",
        "ko": "13.653"
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
        "ok": "60000",
        "ko": "0"
    },
    "minResponseTime": {
        "total": "454",
        "ok": "454",
        "ko": "-"
    },
    "maxResponseTime": {
        "total": "33109",
        "ok": "33109",
        "ko": "-"
    },
    "meanResponseTime": {
        "total": "25034",
        "ok": "25034",
        "ko": "-"
    },
    "standardDeviation": {
        "total": "3620",
        "ok": "3620",
        "ko": "-"
    },
    "percentiles1": {
        "total": "25100",
        "ok": "25100",
        "ko": "-"
    },
    "percentiles2": {
        "total": "27958",
        "ok": "27958",
        "ko": "-"
    },
    "percentiles3": {
        "total": "30871",
        "ok": "30871",
        "ko": "-"
    },
    "percentiles4": {
        "total": "31763",
        "ok": "31763",
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
    "count": 59978,
    "percentage": 100
},
    "group4": {
    "name": "failed",
    "count": 0,
    "percentage": 0
},
    "meanNumberOfRequestsPerSecond": {
        "total": "1224.49",
        "ok": "1224.49",
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
        "total": "60000",
        "ok": "59331",
        "ko": "669"
    },
    "minResponseTime": {
        "total": "13669",
        "ok": "13669",
        "ko": "15238"
    },
    "maxResponseTime": {
        "total": "33127",
        "ok": "33127",
        "ko": "27651"
    },
    "meanResponseTime": {
        "total": "24979",
        "ok": "25003",
        "ko": "22815"
    },
    "standardDeviation": {
        "total": "3496",
        "ok": "3495",
        "ko": "2815"
    },
    "percentiles1": {
        "total": "25075",
        "ok": "25097",
        "ko": "23606"
    },
    "percentiles2": {
        "total": "27919",
        "ok": "27951",
        "ko": "24722"
    },
    "percentiles3": {
        "total": "29773",
        "ok": "29781",
        "ko": "27357"
    },
    "percentiles4": {
        "total": "31677",
        "ok": "31685",
        "ko": "27525"
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
    "count": 59331,
    "percentage": 99
},
    "group4": {
    "name": "failed",
    "count": 669,
    "percentage": 1
},
    "meanNumberOfRequestsPerSecond": {
        "total": "1224.49",
        "ok": "1210.837",
        "ko": "13.653"
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
