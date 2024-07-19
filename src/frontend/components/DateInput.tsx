import React, { useState } from 'react';
import { View, TextInput, StyleSheet, Dimensions } from 'react-native';
import { FontAwesome } from '@expo/vector-icons';

const window = Dimensions.get('window');

const DateInput = ({ placeholder, value, onChangeText }) => {
  const formatValue = (text) => {
    // Formate o valor para incluir as barras automaticamente
    if (text.length === 2 || text.length === 5) {
      return text + '/';
    }
    return text;
  };

  return (
    <View style={styles.container}>
      <FontAwesome name="calendar" size={24} color="#3A8A88" />
      <TextInput
        style={styles.input}
        placeholder={placeholder}
        value={formatValue(value)} // Use o valor formatado
        onChangeText={onChangeText}
        keyboardType="numeric"
        maxLength={8} // Defina o limite máximo de caracteres
      />
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flexDirection: 'row',
    alignItems: 'center',
    borderWidth: 2,
    borderColor: '#3A8A88',
    borderRadius: 8,
    padding: 10,
    marginVertical: 10,
    width: '100%',
    height: window.height * 0.06,
  },
  input: {
    color: "#3A8A88",
    marginLeft: 10,
    flex: 1,
    fontSize: 16,
  },
});

export default DateInput;
