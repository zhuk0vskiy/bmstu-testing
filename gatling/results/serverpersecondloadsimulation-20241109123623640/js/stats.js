var stats = {
    type: "GROUP",
name: "Global Information",
path: "",
pathFormatted: "group_missing-name-b06d1",
stats: {
    "name": "Global Information",
    "numberOfRequests": {
        "total": "120000",
        "ok": "119461",
        "ko": "539"
    },
    "minResponseTime": {
        "total": "754",
        "ok": "754",
        "ko": "10383"
    },
    "maxResponseTime": {
        "total": "36838",
        "ok": "36838",
        "ko": "27028"
    },
    "meanResponseTime": {
        "total": "26245",
        "ok": "26262",
        "ko": "22512"
    },
    "standardDeviation": {
        "total": "3894",
        "ok": "3883",
        "ko": "4581"
    },
    "percentiles1": {
        "total": "26671",
        "ok": "26689",
        "ko": "24202"
    },
    "percentiles2": {
        "total": "28798",
        "ok": "28818",
        "ko": "25442"
    },
    "percentiles3": {
        "total": "32009",
        "ok": "32029",
        "ko": "26560"
    },
    "percentiles4": {
        "total": "33927",
        "ok": "33928",
        "ko": "26832"
    },
    "group1": {
    "name": "t < 800 ms",
    "count": 1,
    "percentage": 0
},
    "group2": {
    "name": "800 ms < t < 1200 ms",
    "count": 21,
    "percentage": 0
},
    "group3": {
    "name": "t > 1200 ms",
    "count": 119439,
    "percentage": 100
},
    "group4": {
    "name": "failed",
    "count": 539,
    "percentage": 0
},
    "meanNumberOfRequestsPerSecond": {
        "total": "2307.692",
        "ok": "2297.327",
        "ko": "10.365"
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
        "ok": "59461",
        "ko": "539"
    },
    "minResponseTime": {
        "total": "754",
        "ok": "754",
        "ko": "10383"
    },
    "maxResponseTime": {
        "total": "35344",
        "ok": "35344",
        "ko": "27028"
    },
    "meanResponseTime": {
        "total": "26295",
        "ok": "26329",
        "ko": "22512"
    },
    "standardDeviation": {
        "total": "3728",
        "ok": "3701",
        "ko": "4581"
    },
    "percentiles1": {
        "total": "26688",
        "ok": "26722",
        "ko": "24202"
    },
    "percentiles2": {
        "total": "28717",
        "ok": "28756",
        "ko": "25442"
    },
    "percentiles3": {
        "total": "31768",
        "ok": "31791",
        "ko": "26560"
    },
    "percentiles4": {
        "total": "33743",
        "ok": "33743",
        "ko": "26832"
    },
    "group1": {
    "name": "t < 800 ms",
    "count": 1,
    "percentage": 0
},
    "group2": {
    "name": "800 ms < t < 1200 ms",
    "count": 21,
    "percentage": 0
},
    "group3": {
    "name": "t > 1200 ms",
    "count": 59439,
    "percentage": 99
},
    "group4": {
    "name": "failed",
    "count": 539,
    "percentage": 1
},
    "meanNumberOfRequestsPerSecond": {
        "total": "1153.846",
        "ok": "1143.481",
        "ko": "10.365"
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
        "total": "9219",
        "ok": "9219",
        "ko": "-"
    },
    "maxResponseTime": {
        "total": "36838",
        "ok": "36838",
        "ko": "-"
    },
    "meanResponseTime": {
        "total": "26196",
        "ok": "26196",
        "ko": "-"
    },
    "standardDeviation": {
        "total": "4054",
        "ok": "4054",
        "ko": "-"
    },
    "percentiles1": {
        "total": "26649",
        "ok": "26648",
        "ko": "-"
    },
    "percentiles2": {
        "total": "28871",
        "ok": "28871",
        "ko": "-"
    },
    "percentiles3": {
        "total": "32306",
        "ok": "32306",
        "ko": "-"
    },
    "percentiles4": {
        "total": "34015",
        "ok": "34015",
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
        "total": "1153.846",
        "ok": "1153.846",
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
