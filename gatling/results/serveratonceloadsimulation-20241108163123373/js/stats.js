var stats = {
    type: "GROUP",
name: "Global Information",
path: "",
pathFormatted: "group_missing-name-b06d1",
stats: {
    "name": "Global Information",
    "numberOfRequests": {
        "total": "100000",
        "ok": "99990",
        "ko": "10"
    },
    "minResponseTime": {
        "total": "450",
        "ok": "450",
        "ko": "18701"
    },
    "maxResponseTime": {
        "total": "28569",
        "ok": "28569",
        "ko": "19722"
    },
    "meanResponseTime": {
        "total": "22871",
        "ok": "22871",
        "ko": "19403"
    },
    "standardDeviation": {
        "total": "4445",
        "ok": "4445",
        "ko": "258"
    },
    "percentiles1": {
        "total": "24289",
        "ok": "24292",
        "ko": "19428"
    },
    "percentiles2": {
        "total": "27075",
        "ok": "27075",
        "ko": "19512"
    },
    "percentiles3": {
        "total": "28071",
        "ok": "28071",
        "ko": "19667"
    },
    "percentiles4": {
        "total": "28346",
        "ok": "28346",
        "ko": "19711"
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
    "count": 99968,
    "percentage": 100
},
    "group4": {
    "name": "failed",
    "count": 10,
    "percentage": 0
},
    "meanNumberOfRequestsPerSecond": {
        "total": "2173.913",
        "ok": "2173.696",
        "ko": "0.217"
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
        "total": "450",
        "ok": "450",
        "ko": "-"
    },
    "maxResponseTime": {
        "total": "28568",
        "ok": "28568",
        "ko": "-"
    },
    "meanResponseTime": {
        "total": "23344",
        "ok": "23344",
        "ko": "-"
    },
    "standardDeviation": {
        "total": "4426",
        "ok": "4426",
        "ko": "-"
    },
    "percentiles1": {
        "total": "25415",
        "ok": "25412",
        "ko": "-"
    },
    "percentiles2": {
        "total": "27288",
        "ok": "27289",
        "ko": "-"
    },
    "percentiles3": {
        "total": "28079",
        "ok": "28079",
        "ko": "-"
    },
    "percentiles4": {
        "total": "28335",
        "ok": "28335",
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
    "count": 49978,
    "percentage": 100
},
    "group4": {
    "name": "failed",
    "count": 0,
    "percentage": 0
},
    "meanNumberOfRequestsPerSecond": {
        "total": "1086.957",
        "ok": "1086.957",
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
        "ok": "49990",
        "ko": "10"
    },
    "minResponseTime": {
        "total": "2247",
        "ok": "2247",
        "ko": "18701"
    },
    "maxResponseTime": {
        "total": "28569",
        "ok": "28569",
        "ko": "19722"
    },
    "meanResponseTime": {
        "total": "22398",
        "ok": "22398",
        "ko": "19403"
    },
    "standardDeviation": {
        "total": "4414",
        "ok": "4414",
        "ko": "258"
    },
    "percentiles1": {
        "total": "22704",
        "ok": "22704",
        "ko": "19428"
    },
    "percentiles2": {
        "total": "26572",
        "ok": "26572",
        "ko": "19512"
    },
    "percentiles3": {
        "total": "28050",
        "ok": "28050",
        "ko": "19667"
    },
    "percentiles4": {
        "total": "28361",
        "ok": "28361",
        "ko": "19711"
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
    "count": 49990,
    "percentage": 100
},
    "group4": {
    "name": "failed",
    "count": 10,
    "percentage": 0
},
    "meanNumberOfRequestsPerSecond": {
        "total": "1086.957",
        "ok": "1086.739",
        "ko": "0.217"
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
