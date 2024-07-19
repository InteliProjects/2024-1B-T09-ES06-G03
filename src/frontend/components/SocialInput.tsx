import React from 'react';
import { View, Text, TextInput, StyleSheet, Dimensions, TouchableOpacity } from 'react-native';
import Svg, { Circle, Path, G, LinearGradient, Stop } from 'react-native-svg';

const window = Dimensions.get('window');

const SocialInput = ({ type = 'Instagram', placeholder }) => {
  const labelStyle = type === 'LinkedIn' ? styles.labelLinkedIn : styles.labelInstagram;
  const Icon = type === 'LinkedIn' ? LinkedInIcon : InstagramIcon;

  return (
    <View style={styles.container}>
      <View style={styles.leftContainer}>
        <Icon />
        <Text style={labelStyle}>{type}</Text>
      </View>
      <TextInput style={[styles.input]} placeholder={placeholder} />
    </View>
  );
};

const LinkedInIcon = () => (
  <TouchableOpacity>
    <Svg width="30" height="30" viewBox="0 0 25 25" fill="none">
      <Circle cx="12.4031" cy="12.4031" r="12.4031" fill="#0076B2" />
      <Path d="M5.6862 10.0737H8.56529V19.349H5.6862V10.0737ZM7.12654 5.45752C7.45676 5.45752 7.77956 5.55558 8.0541 5.73931C8.32864 5.92303 8.5426 6.18416 8.66889 6.48965C8.79519 6.79515 8.82815 7.13129 8.76362 7.45554C8.69908 7.7798 8.53994 8.07761 8.30633 8.31129C8.07272 8.54497 7.77513 8.70403 7.45122 8.76834C7.12731 8.83265 6.79163 8.79932 6.48664 8.67258C6.18165 8.54583 5.92105 8.33136 5.73783 8.05629C5.5546 7.78123 5.45696 7.45793 5.45728 7.1273C5.4577 6.6843 5.63375 6.25959 5.94675 5.94649C6.25976 5.63339 6.6841 5.45752 7.12654 5.45752ZM10.3713 10.0737H13.1311V11.3471H13.1693C13.554 10.6181 14.492 9.84924 15.8926 9.84924C18.8082 9.84288 19.3487 11.7642 19.3487 14.2553V19.349H16.4697V14.8363C16.4697 13.7618 16.4506 12.3786 14.9737 12.3786C13.4968 12.3786 13.2456 13.5501 13.2456 14.7663V19.349H10.3713V10.0737Z" fill="white" />
    </Svg>
  </TouchableOpacity>
);

const InstagramIcon = () => (
  <TouchableOpacity>
    <Svg height="30" width="30" viewBox="0 0 128 128">
      <LinearGradient
        id="a"
        gradientTransform="matrix(1 0 0 -1 594 633)"
        gradientUnits="userSpaceOnUse"
        x1="-566.711"
        x2="-493.288"
        y1="516.569"
        y2="621.43"
      >
        <Stop offset="0" stopColor="#ffb900" />
        <Stop offset="1" stopColor="#9100eb" />
      </LinearGradient>
      <Circle cx="64" cy="64" fill="url(#a)" r="64" />
      <G fill="#fff">
        <Path d="M82.333 104H45.667C33.72 104 24 94.281 24 82.333V45.667C24 33.719 33.72 24 45.667 24h36.666C93.281 24 103 33.719 103 45.667v36.666C103 94.281 93.281 104 82.333 104zM45.667 30.667c-8.271 0-15 6.729-15 15v36.667c0 8.271 6.729 15 15 15h36.666c8.271 0 15-6.729 15-15V45.667c0-8.271-6.729-15-15-15z" />
        <Path d="M64 84c-11.028 0-20-8.973-20-20 0-11.029 8.972-20 20-20s20 8.971 20 20c0 11.027-8.972 20-20 20zm0-33.333c-7.352 0-13.333 5.981-13.333 13.333 0 7.353 5.981 13.333 13.333 13.333s13.333-5.98 13.333-13.333c0-7.352-5.98-13.333-13.333-13.333z" />
        <Circle cx="85.25" cy="42.75" r="4.583" />
      </G>
    </Svg>
  </TouchableOpacity>
);

const styles = StyleSheet.create({
  container: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    padding: 10
  },
  leftContainer: {
    flexDirection: 'row',
    alignItems: 'center',
  },
  labelInstagram: {
    fontSize: 16,
    marginLeft: 10
  },
  labelLinkedIn:{
    fontSize: 16,
    marginLeft: 10,
  },
  input: {
    width: '59%',
    height: window.height * 0.06,
    borderWidth: 2,
    borderRadius: 10,
    paddingLeft: 10,
    borderColor: '#3A8A88',
    marginLeft: 10,
  },
});

export default SocialInput;
