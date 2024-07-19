import React, { useState } from 'react';
import { View, Text, TouchableOpacity, StyleSheet, Dimensions } from 'react-native';
import Icon from 'react-native-vector-icons/Ionicons';

const window = Dimensions.get('window');

const Dropdown = ({ options, selectedValue, onValueChange, placeholder, label, type }: {
  options: { value: any, label: string }[],
  selectedValue: any | null,
  onValueChange: (value: any) => void,
  placeholder: string,
  label: string,
  type: string
}) => {
  const [isOpen, setIsOpen] = useState(false);

  const handleSelect = (value: any) => {
    onValueChange(value);
    setIsOpen(false);
  };

  const dropdownStyle = type === 'grayBlack' ? styles.dropdownGrayBlack : styles.dropdown;
  const selectedValueStyle = type === 'grayBlack' ? styles.selectedTextGrayBlack : styles.selectedText;
  const selectStyle = type === 'grayBlack' ? styles.selectGrayBlack : styles.select;
  const iconColor = type === 'grayBlack' ? '#4A4444' : '#3A8A88';

  return (
    <View style={styles.container}>
      <Text style={styles.label}>{label}</Text>
      <TouchableOpacity
        style={selectStyle}
        onPress={() => setIsOpen(!isOpen)}
      >
        <Text style={selectedValueStyle}>
          {selectedValue ? options.find((option: { value: any, label: string }) => option.value == selectedValue).label : placeholder}
        </Text>
        <Icon name={isOpen ? "chevron-up-outline" : "chevron-down-outline"} size={window.height * 0.025} color={iconColor}/>
      </TouchableOpacity>
      {isOpen && (
        <View style={[dropdownStyle, options.length > 5 ? styles.scrollableDropdown : {}]}>
          {options.map((item) => (
            <TouchableOpacity
              key={item.value}
              style={styles.option}
              onPress={() => handleSelect(item.value)}
            >
              <Text style={styles.optionText}>{item.label}</Text>
            </TouchableOpacity>
          ))}
        </View>
      )}
    </View>
  );
};

const styles = StyleSheet.create({
  selectedTextGrayBlack: {
    color: '#4A4444',
    fontWeight: 'bold',
    textAlign: 'center',
  },
  selectGrayBlack: {
    width: window.width * 0.9,
    height: window.height * 0.06,
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    alignContent: 'center',
    borderWidth: 2,
    borderColor: '#BEBCBC',
    borderRadius: 8,
    padding: 12,
    backgroundColor: '#BEBCBC',
  },
  label: {
    fontSize: 14,
    fontWeight: 'bold',
    marginBottom: 5,
  },
  container: {
    marginTop: window.height * 0.01,
    marginBottom: window.height * 0.01,
  },
  select: {
    width: window.width * 0.9,
    height: window.height * 0.06,
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    borderWidth: 2,
    borderColor: '#3A8A88',
    borderRadius: 8,
    padding: 12,
  },
  selectedText: {
    color: '#3A8A88',
    textAlign: 'center',
  },
  dropdown: {
    position: 'absolute',
    top: window.height * 0.1,
    width: window.width * 0.9,
    borderWidth: 2,
    borderColor: '#3A8A88',
    borderRadius: 8,
    backgroundColor: '#F6F6F6',
    maxHeight: window.height * 0.25,
    zIndex: 3,
  },
  dropdownGrayBlack: {
    position: 'absolute',
    top: window.height * 0.1,
    width: window.width * 0.9,
    borderRadius: 8,
    backgroundColor: '#BEBCBC',
    maxHeight: window.height * 0.25,
    zIndex: 3,
  },
  scrollableDropdown: {
    maxHeight: window.height * 0.5,
  },
  option: {
    padding: 12,
    borderBottomWidth: 1,
    borderBottomColor: '#e2e8f0',
    textAlign: 'center',
  },
  optionText: {
    color: '#000',
    textAlign: 'center',
  },
});

export default Dropdown;
