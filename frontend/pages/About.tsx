import React from 'react';
import styles from './styles';
import axios from 'axios';
import { Link } from "react-router-dom";

interface ChartItemProps {
  department: string;
}

const ChartItem: React.FC<ChartItemProps> = ({ department }) => {
  const handleClick = async () => {
    try {
      await axios.post('/departmentprofit', { department });
      updateButton();
    } catch (error) {
      console.error('Failed to call Go function:', error);
    }
  };

  const updateButton = (content: string, button: any) => {
    // Update the button content
    button.innerHTML = content;
  };

  return (
    <div
      style={styles.chartItem}
      className="chart-item"
      onClick={handleClick}
    >
      {department}
    </div>
  );
};

const About: React.FC = () => {
  const departments = [
    'Дизайн',
    'Печатная',
    'Цех',
    'Материалы',
    'Сольвент',
    'Офсет',
    'POS',
    'Мерч',
    'РИК',
    'По компании',
  ];

  return (
    <div style={styles.containerItem}>
      <Link to="/about">About page</Link>
      <div style={styles.hederItem}>
        Чистая прибыль
        <br />
        в отделах
      </div>
      <div className="chart-section">
        {departments.map((department, index) => (
          <ChartItem key={index} department={department} />
        ))}
      </div>
      <div style={styles.hederItem}>
        Без учета отдела РИК
      </div>
    </div>
  );
};

export default About;
