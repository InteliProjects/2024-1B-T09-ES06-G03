// components/ActivityCard.js
import React from 'react';
import { View, Text, StyleSheet } from 'react-native';

const ActivityCard = ({ date, title, items }) => {
  return (
    <View style={styles.card}>
      <Text style={styles.date}>{date}</Text>
      <Text style={styles.title}>{title}</Text>
      {items.map((item, index) => (
        <Text key={index} style={styles.item}>{item}</Text>
      ))}
    </View>
  );
};

const styles = StyleSheet.create({
  card: {
    backgroundColor: '#F6F6F6',
    padding: 10,
    marginVertical: 5,
    borderRadius: 5,
    elevation: 1,
  },
  date: {
    fontSize: 12,
    color: '#3A8A88',
    marginBottom: 5,
  },
  title: {
    fontSize: 14,
    fontWeight: 'bold',
    marginBottom: 5,
  },
  item: {
    fontSize: 14,
  },
});

export default ActivityCard;
