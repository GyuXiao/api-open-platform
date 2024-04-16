import {
    ParamsType,
    ProColumns,
} from '@ant-design/pro-components';
import '@umijs/max';
import { ProTable } from '@ant-design/pro-components';
import { Modal } from 'antd';
import React from 'react';
export type Props = {
  columns: ProColumns<API.InterfaceInfo>[];
  onCancel: () => void;
  onSubmit: (value: ParamsType) => Promise<void>;
  visible: boolean;
  // values: Partial<API.RuleListItem>;
};
const CreateModel: React.FC<Props> = (props) => {
  const { visible, columns, onCancel, onSubmit } = props;
  return (
    <Modal visible={visible} footer={null} onCancel={() => onCancel?.()}>
      {/* 创建一个ProTable组件,设定它为表单类型,通过columns属性设置表格的列，提交表单时调用onSubmit函数 */}
      <ProTable
        type="form"
        columns={columns}
        onSubmit={async (value) => {onSubmit?.(value)}}
      />
    </Modal>
  );
};
export default CreateModel;
