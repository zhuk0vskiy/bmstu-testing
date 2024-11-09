var stats = {
    type: "GROUP",
name: "Global Information",
path: "",
pathFormatted: "group_missing-name-b06d1",
stats: {
    "name": "Global Information",
    "numberOfRequests": {
        "total": "120000",
        "ok": "119990",
        "ko": "10"
    },
    "minResponseTime": {
        "total": "560",
        "ok": "560",
        "ko": "17928"
    },
    "maxResponseTime": {
        "total": "37548",
        "ok": "37548",
        "ko": "27112"
    },
    "meanResponseTime": {
        "total": "26995",
        "ok": "26995",
        "ko": "24991"
    },
    "standardDeviation": {
        "total": "4902",
        "ok": "4902",
        "ko": "2814"
    },
    "percentiles1": {
        "total": "27253",
        "ok": "27253",
        "ko": "26369"
    },
    "percentiles2": {
        "total": "30937",
        "ok": "30937",
        "ko": "26753"
    },
    "percentiles3": {
        "total": "34402",
        "ok": "34402",
        "ko": "26958"
    },
    "percentiles4": {
        "total": "34975",
        "ok": "34975",
        "ko": "27081"
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
    "count": 119968,
    "percentage": 100
},
    "group4": {
    "name": "failed",
    "count": 10,
    "percentage": 0
},
    "meanNumberOfRequestsPerSecond": {
        "total": "2264.151",
        "ok": "2263.962",
        "ko": "0.189"
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
        "total": "560",
        "ok": "560",
        "ko": "-"
    },
    "maxResponseTime": {
        "total": "37548",
        "ok": "37548",
        "ko": "-"
    },
    "meanResponseTime": {
        "total": "27043",
        "ok": "27043",
        "ko": "-"
    },
    "standardDeviation": {
        "total": "4856",
        "ok": "4856",
        "ko": "-"
    },
    "percentiles1": {
        "total": "27354",
        "ok": "27341",
        "ko": "-"
    },
    "percentiles2": {
        "total": "30890",
        "ok": "30890",
        "ko": "-"
    },
    "percentiles3": {
        "total": "34423",
        "ok": "34423",
        "ko": "-"
    },
    "percentiles4": {
        "total": "34989",
        "ok": "34989",
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
        "total": "1132.075",
        "ok": "1132.075",
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
        "ok": "59990",
        "ko": "10"
    },
    "minResponseTime": {
        "total": "13544",
        "ok": "13544",
        "ko": "17928"
    },
    "maxResponseTime": {
        "total": "37548",
        "ok": "37548",
        "ko": "27112"
    },
    "meanResponseTime": {
        "total": "26946",
        "ok": "26946",
        "ko": "24991"
    },
    "standardDeviation": {
        "total": "4947",
        "ok": "4948",
        "ko": "2814"
    },
    "percentiles1": {
        "total": "27102",
        "ok": "27104",
        "ko": "26369"
    },
    "percentiles2": {
        "total": "30982",
        "ok": "30983",
        "ko": "26753"
    },
    "percentiles3": {
        "total": "34368",
        "ok": "34368",
        "ko": "26958"
    },
    "percentiles4": {
        "total": "34970",
        "ok": "34970",
        "ko": "27081"
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
    "count": 59990,
    "percentage": 100
},
    "group4": {
    "name": "failed",
    "count": 10,
    "percentage": 0
},
    "meanNumberOfRequestsPerSecond": {
        "total": "1132.075",
        "ok": "1131.887",
        "ko": "0.189"
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
