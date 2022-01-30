// import App from 'next/app'
var AWS = require('aws-sdk');

// Create CloudWatch service object
var cw = new AWS.CloudWatch({apiVersion: '2010-08-01'});

export function reportWebVitals(metric) {
  const name = "mypage"
  var params = {
    MetricData: [
      {
        MetricName: 'PageAnalytics',
        Dimensions: [
          {
            Name: 'NextJSAPP',
            Value: 'Initial'
          },
        ],
        Unit: 'None',
        Value: 0.0
      },
    ],
    Namespace: 'Page/Analytics'
  };
  console.log("=======CW Before=======");
  console.log(params);

  switch (metric.name) {
    case 'FCP':
      console.log("=======FCP=======");
      console.log(metric);
      params["MetricData"][0]["Dimensions"][0]["Value"] = metric.name;
      params["MetricData"][0]["Dimensions"]["Value"] = name;
      params["MetricData"][0]["Value"] = metric["value"];
      break
    case 'LCP':
      console.log("=======LCP=======");
      console.log(metric);
      params["MetricData"][0]["Dimensions"][0]["Value"] = metric.name;
      params["MetricData"][0]["Dimensions"]["Value"] = name;
      params["MetricData"][0]["Value"] = metric["value"];
      break
    case 'CLS':
      console.log("=======CLS=======");
      console.log(metric);
      params["MetricData"][0]["Dimensions"][0]["Value"] = metric.name;
      params["MetricData"][0]["Dimensions"]["Value"] = name;
      params["MetricData"][0]["Value"] = metric["value"];
      break
    case 'FID':
      console.log("=======FID=======");
      console.log(metric);
      params["MetricData"][0]["Dimensions"][0]["Value"] = metric.name;
      params["MetricData"][0]["Dimensions"]["Value"] = name;
      params["MetricData"][0]["Value"] = metric["value"];
      break
    case 'TTFB':
      console.log("=======TTFB=======");
      console.log(metric);
      params["MetricData"][0]["Dimensions"][0]["Value"] = metric.name;
      params["MetricData"][0]["Dimensions"]["Value"] = name;
      params["MetricData"][0]["Value"] = metric["value"];
      break
    default:
      break
  }

  console.log("=======CW After=======");
  console.log(params);
  cw.putMetricData(params, function(err, data) {
    if (err) {
      console.log("Error", err);
    } else {
      console.log("Success", JSON.stringify(data));
    }
  });
}

function MyApp({ Component, pageProps }) {
  return <Component {...pageProps} />
}

export default MyApp