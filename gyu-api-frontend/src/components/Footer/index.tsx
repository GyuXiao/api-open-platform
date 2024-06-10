import { GithubOutlined } from '@ant-design/icons';
import { DefaultFooter } from '@ant-design/pro-components';
import React from 'react';

const Footer: React.FC = () => {
  return (
    <DefaultFooter
      style={{
        background: 'none',
      }}
      links={[
        {
          key: 'Ant Design Pro',
          title: 'Gyu-Api-Platform',
          href: 'https://github.com/GyuXiao/api-open-platform',
          blankTarget: true,
        },
        {
          key: 'github',
          title: <GithubOutlined />,
          href: 'https://github.com/GyuXiao',
          blankTarget: true,
        },
      ]}
    />
  );
};

export default Footer;
