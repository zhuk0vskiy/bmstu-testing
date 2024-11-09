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
        "total": "331",
        "ok": "-",
        "ko": "331"
    },
    "maxResponseTime": {
        "total": "28624",
        "ok": "-",
        "ko": "28624"
    },
    "meanResponseTime": {
        "total": "16182",
        "ok": "-",
        "ko": "16182"
    },
    "standardDeviation": {
        "total": "4025",
        "ok": "-",
        "ko": "4025"
    },
    "percentiles1": {
        "total": "16002",
        "ok": "-",
        "ko": "16001"
    },
    "percentiles2": {
        "total": "18438",
        "ok": "-",
        "ko": "18463"
    },
    "percentiles3": {
        "total": "24132",
        "ok": "-",
        "ko": "24133"
    },
    "percentiles4": {
        "total": "27149",
        "ok": "-",
        "ko": "27149"
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
        "total": "2857.143",
        "ok": "-",
        "ko": "2857.143"
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
        "total": "331",
        "ok": "-",
        "ko": "331"
    },
    "maxResponseTime": {
        "total": "27469",
        "ok": "-",
        "ko": "27469"
    },
    "meanResponseTime": {
        "total": "15602",
        "ok": "-",
        "ko": "15602"
    },
    "standardDeviation": {
        "total": "3759",
        "ok": "-",
        "ko": "3759"
    },
    "percentiles1": {
        "total": "15720",
        "ok": "-",
        "ko": "15720"
    },
    "percentiles2": {
        "total": "17461",
        "ok": "-",
        "ko": "17460"
    },
    "percentiles3": {
        "total": "22198",
        "ok": "-",
        "ko": "22198"
    },
    "percentiles4": {
        "total": "25973",
        "ok": "-",
        "ko": "25973"
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
        "total": "1428.571",
        "ok": "-",
        "ko": "1428.571"
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
        "total": "1721",
        "ok": "-",
        "ko": "1721"
    },
    "maxResponseTime": {
        "total": "28624",
        "ok": "-",
        "ko": "28624"
    },
    "meanResponseTime": {
        "total": "16763",
        "ok": "-",
        "ko": "16763"
    },
    "standardDeviation": {
        "total": "4195",
        "ok": "-",
        "ko": "4195"
    },
    "percentiles1": {
        "total": "16314",
        "ok": "-",
        "ko": "16314"
    },
    "percentiles2": {
        "total": "19066",
        "ok": "-",
        "ko": "19067"
    },
    "percentiles3": {
        "total": "25431",
        "ok": "-",
        "ko": "25430"
    },
    "percentiles4": {
        "total": "27555",
        "ok": "-",
        "ko": "27555"
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
        "total": "1428.571",
        "ok": "-",
        "ko": "1428.571"
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
