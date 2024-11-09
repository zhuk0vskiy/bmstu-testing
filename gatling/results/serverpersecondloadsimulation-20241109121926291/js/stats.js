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
        "total": "256",
        "ok": "-",
        "ko": "256"
    },
    "maxResponseTime": {
        "total": "38415",
        "ok": "-",
        "ko": "38415"
    },
    "meanResponseTime": {
        "total": "21872",
        "ok": "-",
        "ko": "21872"
    },
    "standardDeviation": {
        "total": "8012",
        "ok": "-",
        "ko": "8012"
    },
    "percentiles1": {
        "total": "22158",
        "ok": "-",
        "ko": "22161"
    },
    "percentiles2": {
        "total": "30096",
        "ok": "-",
        "ko": "30098"
    },
    "percentiles3": {
        "total": "32618",
        "ok": "-",
        "ko": "32618"
    },
    "percentiles4": {
        "total": "34893",
        "ok": "-",
        "ko": "34896"
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
        "total": "2264.151",
        "ok": "-",
        "ko": "2264.151"
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
        "total": "256",
        "ok": "-",
        "ko": "256"
    },
    "maxResponseTime": {
        "total": "36472",
        "ok": "-",
        "ko": "36472"
    },
    "meanResponseTime": {
        "total": "19905",
        "ok": "-",
        "ko": "19905"
    },
    "standardDeviation": {
        "total": "8221",
        "ok": "-",
        "ko": "8221"
    },
    "percentiles1": {
        "total": "17170",
        "ok": "-",
        "ko": "17170"
    },
    "percentiles2": {
        "total": "30048",
        "ok": "-",
        "ko": "30048"
    },
    "percentiles3": {
        "total": "31752",
        "ok": "-",
        "ko": "31752"
    },
    "percentiles4": {
        "total": "34577",
        "ok": "-",
        "ko": "34584"
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
        "total": "1132.075",
        "ok": "-",
        "ko": "1132.075"
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
        "total": "4487",
        "ok": "-",
        "ko": "4487"
    },
    "maxResponseTime": {
        "total": "38415",
        "ok": "-",
        "ko": "38415"
    },
    "meanResponseTime": {
        "total": "23838",
        "ok": "-",
        "ko": "23838"
    },
    "standardDeviation": {
        "total": "7283",
        "ok": "-",
        "ko": "7283"
    },
    "percentiles1": {
        "total": "25938",
        "ok": "-",
        "ko": "25940"
    },
    "percentiles2": {
        "total": "30219",
        "ok": "-",
        "ko": "30219"
    },
    "percentiles3": {
        "total": "33278",
        "ok": "-",
        "ko": "33278"
    },
    "percentiles4": {
        "total": "35177",
        "ok": "-",
        "ko": "35178"
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
        "total": "1132.075",
        "ok": "-",
        "ko": "1132.075"
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
