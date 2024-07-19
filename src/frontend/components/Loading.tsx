import React, { useEffect, useRef } from 'react';
import { View, StyleSheet, Animated } from 'react-native';

const useLoopingAnimation = (delay) => {
  const animatedValue = useRef(new Animated.Value(0)).current;

  useEffect(() => {
    const startAnimation = () => {
      Animated.loop(
        Animated.sequence([
          Animated.timing(animatedValue, {
            toValue: 1,
            duration: 500,
            useNativeDriver: false,
          }),
          Animated.timing(animatedValue, {
            toValue: 0,
            duration: 500,
            useNativeDriver: false,
          }),
        ])
      ).start();
    };

    const timer = setTimeout(startAnimation, delay);

    return () => clearTimeout(timer);
  }, [animatedValue, delay]);

  return animatedValue;
};

const AnimatedCircle = ({ style, delay }) => {
  const animatedValue = useLoopingAnimation(delay);

  const top = animatedValue.interpolate({
    inputRange: [0, 0.4, 1],
    outputRange: [60, 0, 0],
  });

  const height = animatedValue.interpolate({
    inputRange: [0, 0.4, 1],
    outputRange: [5, 20, 20],
  });

  const borderRadius = animatedValue.interpolate({
    inputRange: [0, 0.4, 1],
    outputRange: [25, 50, 50],
  });

  const scaleX = animatedValue.interpolate({
    inputRange: [0, 0.4, 1],
    outputRange: [1.7, 1, 1],
  });

  return (
    <Animated.View
      style={[
        style,
        {
          top,
          height,
          borderRadius,
          transform: [{ scaleX }],
        },
      ]}
    />
  );
};

const AnimatedShadow = ({ style, delay }) => {
  const animatedValue = useLoopingAnimation(delay);

  const scaleX = animatedValue.interpolate({
    inputRange: [0, 0.4, 1],
    outputRange: [1.5, 1, 0.2],
  });

  const opacity = animatedValue.interpolate({
    inputRange: [0, 0.4, 1],
    outputRange: [1, 0.7, 0.4],
  });

  return (
    <Animated.View
      style={[
        style,
        {
          transform: [{ scaleX }],
          opacity,
        },
      ]}
    />
  );
};

const Loading = () => {
  return (
    <View style={styles.wrapper}>
      <AnimatedCircle style={[styles.circle, { left: '15%' }]} delay={0} />
      <AnimatedCircle style={[styles.circle, { left: '45%' }]} delay={300} />
      <AnimatedCircle style={[styles.circle, { right: '15%' }]} delay={600} />
      <AnimatedShadow style={[styles.shadow, { left: '15%' }]} delay={0} />
      <AnimatedShadow style={[styles.shadow, { left: '45%' }]} delay={300} />
      <AnimatedShadow style={[styles.shadow, { right: '15%' }]} delay={600} />
    </View>
  );
};

const styles = StyleSheet.create({
  wrapper: {
    width: 200,
    height: 60,
    position: 'relative',
    zIndex: 1,
  },
  circle: {
    width: 20,
    height: 20,
    position: 'absolute',
    borderRadius: 50,
    borderWidth: 0.2,
    borderColor: 'black',
    backgroundColor: '#fff',
    transformOrigin: '50%',
  },
  shadow: {
    width: 20,
    height: 4,
    borderRadius: 50,
    backgroundColor: 'rgba(0,0,0,0.9)',
    position: 'absolute',
    top: 62,
    transformOrigin: '50%',
    zIndex: -1,
  },
});

export default Loading;
