var stats = {
    type: "GROUP",
name: "Global Information",
path: "",
pathFormatted: "group_missing-name-b06d1",
stats: {
    "name": "Global Information",
    "numberOfRequests": {
        "total": "120000",
        "ok": "119374",
        "ko": "626"
    },
    "minResponseTime": {
        "total": "4722",
        "ok": "4722",
        "ko": "15469"
    },
    "maxResponseTime": {
        "total": "37428",
        "ok": "37428",
        "ko": "29585"
    },
    "meanResponseTime": {
        "total": "26602",
        "ok": "26611",
        "ko": "24836"
    },
    "standardDeviation": {
        "total": "4007",
        "ok": "4010",
        "ko": "2899"
    },
    "percentiles1": {
        "total": "25895",
        "ok": "25904",
        "ko": "24558"
    },
    "percentiles2": {
        "total": "29528",
        "ok": "29555",
        "ko": "27506"
    },
    "percentiles3": {
        "total": "33893",
        "ok": "33898",
        "ko": "28853"
    },
    "percentiles4": {
        "total": "36372",
        "ok": "36374",
        "ko": "29221"
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
    "count": 119374,
    "percentage": 99
},
    "group4": {
    "name": "failed",
    "count": 626,
    "percentage": 1
},
    "meanNumberOfRequestsPerSecond": {
        "total": "2222.222",
        "ok": "2210.63",
        "ko": "11.593"
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
        "ok": "59374",
        "ko": "626"
    },
    "minResponseTime": {
        "total": "4722",
        "ok": "4722",
        "ko": "15469"
    },
    "maxResponseTime": {
        "total": "37427",
        "ok": "37427",
        "ko": "29585"
    },
    "meanResponseTime": {
        "total": "26542",
        "ok": "26560",
        "ko": "24836"
    },
    "standardDeviation": {
        "total": "4012",
        "ok": "4018",
        "ko": "2899"
    },
    "percentiles1": {
        "total": "25749",
        "ok": "25766",
        "ko": "24558"
    },
    "percentiles2": {
        "total": "29364",
        "ok": "29409",
        "ko": "27506"
    },
    "percentiles3": {
        "total": "33873",
        "ok": "33884",
        "ko": "28853"
    },
    "percentiles4": {
        "total": "36562",
        "ok": "36565",
        "ko": "29221"
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
    "count": 59374,
    "percentage": 99
},
    "group4": {
    "name": "failed",
    "count": 626,
    "percentage": 1
},
    "meanNumberOfRequestsPerSecond": {
        "total": "1111.111",
        "ok": "1099.519",
        "ko": "11.593"
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
        "total": "11405",
        "ok": "11405",
        "ko": "-"
    },
    "maxResponseTime": {
        "total": "37428",
        "ok": "37428",
        "ko": "-"
    },
    "meanResponseTime": {
        "total": "26661",
        "ok": "26661",
        "ko": "-"
    },
    "standardDeviation": {
        "total": "4002",
        "ok": "4002",
        "ko": "-"
    },
    "percentiles1": {
        "total": "26071",
        "ok": "26083",
        "ko": "-"
    },
    "percentiles2": {
        "total": "29687",
        "ok": "29687",
        "ko": "-"
    },
    "percentiles3": {
        "total": "33935",
        "ok": "33950",
        "ko": "-"
    },
    "percentiles4": {
        "total": "36308",
        "ok": "36308",
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
        "total": "1111.111",
        "ok": "1111.111",
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
