import React, { useState, useEffect } from 'react';
import { TextInput, View, Text, StyleSheet, Dimensions, Image, TouchableOpacity } from 'react-native';

const window = Dimensions.get('window');

const InputComponent = ({ type, value, placeholder, maxLength, label, onChangeText }) => {
  const [inputValue, setInputValue] = useState(value || '');
  const [isFocused, setIsFocused] = useState(false);
  const [isPasswordVisible, setIsPasswordVisible] = useState(false);
  const isTextArea = type === 'textarea';

  useEffect(() => {
    setInputValue(value || '');
  }, [value]);

  const togglePasswordVisibility = () => {
    setIsPasswordVisible(!isPasswordVisible);
  };

  const handleInputChange = (text) => {
    setInputValue(text);
    if (onChangeText) {
      onChangeText(text);
    }
  };

  return (
    <View style={styles.container}>
      {label && <Text style={styles.label}>{label}</Text>}
      <View style={styles.inputWrapper}>
        <TextInput
          style={isTextArea ? styles.textArea : styles.input}
          value={inputValue}
          onChangeText={handleInputChange}
          multiline={isTextArea}
          maxLength={maxLength}
          onFocus={() => setIsFocused(true)}
          onBlur={() => setIsFocused(false)}
          textAlignVertical={isTextArea ? 'top' : 'center'}
          secureTextEntry={type === 'password' && !isPasswordVisible}
          placeholder={placeholder}
          placeholderTextColor="#3A8A88"
        />
        {type === 'password' && (
          <TouchableOpacity onPress={togglePasswordVisibility} style={styles.eyeIconWrapper}>
            <Image source={require('../assets/eyeIcon.png')} style={styles.eyeIcon} />
          </TouchableOpacity>
        )}
      </View>
      {isTextArea && (
        <Text style={styles.counter}>{maxLength - inputValue.length} restantes</Text>
      )}
    </View>
  );
};

const styles = StyleSheet.create({
  label: {
    fontSize: 14,
    fontWeight: 'bold',
    marginBottom: 5,
  },
  container: {
    marginTop: window.height * 0.01,
    marginBottom: window.height * 0.01,
  },
  inputWrapper: {
    flexDirection: 'row',
    alignItems: 'center',
    borderColor: '#3A8A88',
    borderWidth: 2,
    borderRadius: 8,
    width: window.width * 0.9,
  },
  input: {
    flex: 1,
    height: window.height * 0.06,
    paddingLeft: 10,
  },
  textArea: {
    width: '100%',
    height: window.height * 0.2,
    paddingLeft: 10,
    paddingTop: 10,
    paddingBottom: 20,
    textAlignVertical: 'top',
  },
  eyeIconWrapper: {
    padding: 10,
  },
  eyeIcon: {
    width: window.height * 0.024,
    height: window.height * 0.016,
    tintColor: '#3A8A88',
  },
  counter: {
    position: 'absolute',
    bottom: 5,
    right: 10,
    color: '#3A8A88',
  },
});

export default InputComponent;
