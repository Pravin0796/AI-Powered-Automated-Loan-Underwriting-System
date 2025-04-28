import React from 'react';
import { Card, Typography, Tooltip, Button, Row, Col } from 'antd';
import { InfoCircleOutlined, FileSearchOutlined, CreditCardOutlined } from '@ant-design/icons';

const { Title, Text } = Typography;

const ViewCreditPage: React.FC = () => {
  const credits = [
    {
      id: 1,
      name: 'Personal Loan',
      amount: '$15,000',
      status: 'Approved',
    },
    {
      id: 2,
      name: 'Home Loan',
      amount: '$250,000',
      status: 'Pending',
    },
    {
      id: 3,
      name: 'Auto Loan',
      amount: '$20,000',
      status: 'Rejected',
    },
  ];

  return (
    <div style={{ padding: '24px' }}>
      <Title level={2}>View Credit</Title>
      <Row gutter={[16, 16]}>
        {credits.map((credit) => (
          <Col xs={24} sm={12} md={8} key={credit.id}>
            <Card
              hoverable
              style={{
                borderRadius: '8px',
                transition: '0.3s',
              }}
            >
              <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
                <Text strong>{credit.name}</Text>
                <Tooltip title="More Info">
                  <InfoCircleOutlined style={{ fontSize: '16px', cursor: 'pointer' }} />
                </Tooltip>
              </div>
              <Text type="secondary">Amount: {credit.amount}</Text>
              <div style={{ marginTop: '8px' }}>
                <Text>
                  Status:{' '}
                  <strong
                    style={{
                      color:
                        credit.status === 'Approved'
                          ? 'green'
                          : credit.status === 'Pending'
                          ? 'orange'
                          : 'red',
                    }}
                  >
                    {credit.status}
                  </strong>
                </Text>
              </div>
              <div style={{ display: 'flex', justifyContent: 'space-between', marginTop: '16px' }}>
                <Button type="primary" icon={<FileSearchOutlined />}>
                  Details
                </Button>
                <Button type="default" icon={<CreditCardOutlined />}>
                  Apply Again
                </Button>
              </div>
            </Card>
          </Col>
        ))}
      </Row>
    </div>
  );
};

export default ViewCreditPage;