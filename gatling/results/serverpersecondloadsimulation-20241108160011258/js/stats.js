var stats = {
    type: "GROUP",
name: "Global Information",
path: "",
pathFormatted: "group_missing-name-b06d1",
stats: {
    "name": "Global Information",
    "numberOfRequests": {
        "total": "120000",
        "ok": "120000",
        "ko": "0"
    },
    "minResponseTime": {
        "total": "863",
        "ok": "863",
        "ko": "-"
    },
    "maxResponseTime": {
        "total": "36811",
        "ok": "36811",
        "ko": "-"
    },
    "meanResponseTime": {
        "total": "28020",
        "ok": "28020",
        "ko": "-"
    },
    "standardDeviation": {
        "total": "3092",
        "ok": "3092",
        "ko": "-"
    },
    "percentiles1": {
        "total": "27520",
        "ok": "27520",
        "ko": "-"
    },
    "percentiles2": {
        "total": "30254",
        "ok": "30253",
        "ko": "-"
    },
    "percentiles3": {
        "total": "33062",
        "ok": "33062",
        "ko": "-"
    },
    "percentiles4": {
        "total": "35255",
        "ok": "35255",
        "ko": "-"
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
    "count": 119978,
    "percentage": 100
},
    "group4": {
    "name": "failed",
    "count": 0,
    "percentage": 0
},
    "meanNumberOfRequestsPerSecond": {
        "total": "2181.818",
        "ok": "2181.818",
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
        "total": "60000",
        "ok": "60000",
        "ko": "0"
    },
    "minResponseTime": {
        "total": "863",
        "ok": "863",
        "ko": "-"
    },
    "maxResponseTime": {
        "total": "36811",
        "ok": "36811",
        "ko": "-"
    },
    "meanResponseTime": {
        "total": "27883",
        "ok": "27883",
        "ko": "-"
    },
    "standardDeviation": {
        "total": "3134",
        "ok": "3134",
        "ko": "-"
    },
    "percentiles1": {
        "total": "27388",
        "ok": "27388",
        "ko": "-"
    },
    "percentiles2": {
        "total": "30134",
        "ok": "30134",
        "ko": "-"
    },
    "percentiles3": {
        "total": "32967",
        "ok": "32967",
        "ko": "-"
    },
    "percentiles4": {
        "total": "35143",
        "ok": "35143",
        "ko": "-"
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
    "count": 59978,
    "percentage": 100
},
    "group4": {
    "name": "failed",
    "count": 0,
    "percentage": 0
},
    "meanNumberOfRequestsPerSecond": {
        "total": "1090.909",
        "ok": "1090.909",
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
        "ok": "60000",
        "ko": "0"
    },
    "minResponseTime": {
        "total": "13980",
        "ok": "13980",
        "ko": "-"
    },
    "maxResponseTime": {
        "total": "36811",
        "ok": "36811",
        "ko": "-"
    },
    "meanResponseTime": {
        "total": "28158",
        "ok": "28158",
        "ko": "-"
    },
    "standardDeviation": {
        "total": "3044",
        "ok": "3044",
        "ko": "-"
    },
    "percentiles1": {
        "total": "27692",
        "ok": "27690",
        "ko": "-"
    },
    "percentiles2": {
        "total": "30396",
        "ok": "30407",
        "ko": "-"
    },
    "percentiles3": {
        "total": "33138",
        "ok": "33138",
        "ko": "-"
    },
    "percentiles4": {
        "total": "35421",
        "ok": "35421",
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
        "total": "1090.909",
        "ok": "1090.909",
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
