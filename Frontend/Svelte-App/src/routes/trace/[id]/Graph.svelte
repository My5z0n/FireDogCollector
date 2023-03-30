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
  import ChartDataLabels from "chartjs-plugin-datalabels";
  Chart.register(ChartDataLabels);
  Chart.defaults.plugins.legend.display = false;
  Chart.defaults.plugins.datalabels.align = "end";

  import { onMount } from "svelte";

  export let SpanList = [];

  let GraphType = "dendogram";
  let mountedCanvas = undefined;

  export let NewSpanList = [];
  export let mainChart;

  let testAl = 0;
  let chartCanvas;
  let config = {
    type: GraphType,
    data: {
      labels: NewSpanList.map((obj) => obj.name),
      datasets: [
        {
          pointBackgroundColor: "steelblue",
          pointRadius: 8,
          pointHoverRadius: 10,
          data: NewSpanList,

          clip: { left: 125, top: 0, right: 125, bottom: 0 },
        },
      ],
    },
    options: {
      layout: {
        padding: {
          right: 100,
          left: 100,
        },
      },
      plugins: {
        colors: {
          forceOverride: true,
        },
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
        datalabels: {
          anchor: "end",
          font: {
            size: 10,
            weight: "bold",
          },
          formatter: function (value, context) {
            return context.chart.data.labels[context.dataIndex];
          },
          offset: function (context) {
            let test = 0;
            test = context.dataIndex;
            return test % 2 === 0 ? -40 : 0;
          },
        },
      },
    },
  };

  onMount(() => {
    mountedCanvas = chartCanvas.getContext("2d");
  });

  $: if (mountedCanvas && SpanList.length > 0) {
    NewSpanList = SpanList.map((obj) => {
      let par = 0;
      let p_id = obj.parent_span_id;
      if (p_id != "") {
        par = SpanList.findIndex((x) => x.span_id == p_id);
        return { name: obj.span_name, parent: par, ...obj };
      } else {
        return { name: obj.span_name, ...obj };
      }
    });
    config.data.datasets[0].data = NewSpanList;
    config.data.datasets[0].pointBackgroundColor = NewSpanList.map((obj) =>
      obj.AnomalyDetected ? "red" : "steelblue"
    );
    config.data.labels = NewSpanList.map((obj) => obj.name);
    mainChart = new Chart(mountedCanvas, config);
  }
</script>

<div>
  <canvas bind:this={chartCanvas} />
</div>
