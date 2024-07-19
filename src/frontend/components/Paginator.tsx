import React from 'react';
import { View, Animated, useWindowDimensions } from 'react-native';

type PaginatorProps = {
    data: { id: number; text: string; image: any }[];
    scrollX: Animated.Value;
};

// Paginação para o carrosel
const Paginator: React.FC<PaginatorProps> = ({ data, scrollX }) => {
    const { width } = useWindowDimensions();

    return (
        <View className="flex flex-row h-[20px] items-center">
            {data.map((_, i) => {
                const inputRange = [(i - 1) * width, i * width, (i + 1) * width];
                const dotWidth = scrollX.interpolate({
                    inputRange,
                    outputRange: [10, 20, 10],
                    extrapolate: 'clamp',
                });

                const opacity = scrollX.interpolate({
                    inputRange,
                    outputRange: [0.3, 1, 0.3],
                    extrapolate: 'clamp',
                });

                return (
                    <Animated.View className="h-[10px] rounded-full bg-green-10 mx-[8px]" style={{ width: dotWidth, opacity }} key={i.toString()} />
                );
            })}
        </View>
    );
};

export default Paginator;
