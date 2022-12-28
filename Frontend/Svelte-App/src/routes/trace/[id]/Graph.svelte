<script>
    import { Chart } from "chart.js/auto";
    import {
        DendogramChart,
        DendogramController,
        DendrogramChart,
        DendrogramController,
        EdgeLine,
        ForceDirectedGraphChart,
        ForceDirectedGraphController,
        GraphChart,
        GraphController,
        TreeChart,
        TreeController,
    } from "chartjs-chart-graph";
    Chart.register(
        DendogramChart,
        DendogramController,
        DendrogramChart,
        DendrogramController,
        EdgeLine,
        ForceDirectedGraphChart,
        ForceDirectedGraphController,
        GraphChart,
        GraphController,
        TreeChart,
        TreeController
    );
    Chart.defaults.plugins.legend.display = false;

    import { onMount } from "svelte";

    export let SpanList = [];

    let type = "dendogram";
    let pg = undefined;

    export let NewSpanList = [
        { name: "1" },
        { name: "2", parent: 0 },
        { name: "3", parent: 1 },
        { name: "4", parent: 2 },
        { name: "5", parent: 2 },
        { name: "6", parent: 3 },
        { name: "7", parent: 4 },
        { name: "8", parent: 5 },
    ];
    export let myChart;
    let chartCanvas;
    let config = {
        type,
        data: {
            labels: NewSpanList.map((obj) => obj.name),
            datasets: [
                {
                    pointBackgroundColor: "steelblue",
                    pointRadius: 8,
                    pointHoverRadius: 10,
                    data: NewSpanList,
                    pointHoverBorderColor: "Navy",
                    //pointBackgroundColor: ["black", "rgb(255, 99, 132)"],
                },
            ],
        },
        options: {
            plugins: {
                tooltip: {
                    displayColors: false,
                    callbacks: {
                        label: function (context) {
                            let t = context;
                            return t.raw.span_id || "Root";
                        },
                        title: function (context) {
                            let t = context[0];
                            return t.raw.span_name;
                        },
                    },
                },
            },
        },
    };

    onMount(() => {
        console.log("MMMM update!");
        pg = chartCanvas.getContext("2d");
    });

    $: if (pg && SpanList.length > 0) {
        let test = [];
        NewSpanList = SpanList.map((obj) => {
            let par = 0;
            let p_id = obj.parent_span_id;
            if (p_id != "") {
                par = SpanList.findIndex((x) => x.span_id == p_id);
                return { name: obj.span_name, parent: par, ...obj };
            } else {
                test = { name: obj.span_name, ...obj };
                return test;
            }
        });
        console.log("New update!");
        console.log(NewSpanList);
        config.data.datasets[0].data = NewSpanList;
        config.data.labels = NewSpanList.map((obj) => obj.name);
        myChart = new Chart(pg, config);
    }
</script>

<div>
    <canvas bind:this={chartCanvas} />
</div>
