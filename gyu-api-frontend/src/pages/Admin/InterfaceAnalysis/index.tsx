import { listTopInvokeInterfaceInfoUsingGET } from '@/services/gyuapi-backend/analysisController';
import { PageContainer } from '@ant-design/pro-components';
import '@umijs/max';
import { Input, message } from 'antd';
import ReactECharts from 'echarts-for-react';
import React, { useEffect, useState } from 'react';

const InterfaceAnalysis: React.FC = () => {
  const [data, setData] = useState<API.InvokeInterface[]>([]);
  const [loading, setLoading] = useState(true);
  const [topN, setTopN] = useState(3);

  const loadData = async () => {
    setLoading(true);
    try {
      // 获取数据
      const res = await listTopInvokeInterfaceInfoUsingGET({ limit: topN });
      if (res?.data) {
        // @ts-ignore
        setData(res?.data.records);
      } else {
        setData([]);
      }
    } catch (error: any) {
      // 请求失败时提示错误信息
      message.error('请求失败，' + error.message);
    }

    setLoading(false);
  };

  useEffect(() => {
    loadData();
  }, [topN]);

  const chartData = data.map((item) => {
    return {
      value: item.totalNum,
      name: item.name,
    };
  });

  const option = {
    title: {
      text: '调用次数最多的接口 TopN',
      left: 'center',
    },
    tooltip: {
      trigger: 'item',
    },
    legend: {
      bottom: '5%',
      left: 'center',
    },
    series: [
      {
        name: '接口名称 && 调用次数',
        type: 'pie',
        radius: ['30%', '60%'],
        avoidLabelOverlap: false,
        label: {
          show: false,
          position: 'center',
        },
        emphasis: {
          label: {
            show: true,
            fontSize: 15,
            fontWeight: 'bold',
          },
        },
        labelLine: {
          show: false,
        },
        data: chartData,
      },
    ],
  };

  return (
    <PageContainer>
      <div
        style={{
          textAlign: 'left',
          margin: 'auto',
          marginBottom: '20px',
        }}
      >
        <b>TopN : </b>
        <Input
          style={{ width: 200 }}
          type="number"
          placeholder="Enter TopN value"
          value={topN}
          onChange={(e) => setTopN(e.target.value)}
        />
      </div>
      {/* 使用 ReactECharts 组件，传入图表配置 */}
      <ReactECharts
        loadingOption={{
          showLoading: loading,
        }}
        option={option}
      />
    </PageContainer>
  );
};
export default InterfaceAnalysis;
