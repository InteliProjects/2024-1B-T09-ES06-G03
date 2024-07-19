//Exemplo de uso: <CeoAvatar size={'w-10 h-10'} link={""} />

import React from 'react';
import { View, Image, ImageStyle } from 'react-native';

interface CeoAvatarProps {
  size: string;
  link: string;
}

export default function CeoAvatar({ size, link }: CeoAvatarProps) {
  return (
    <View>
      <Image
        source={link ? { uri: link } : require('../assets/AvatarPlaceholder.jpg')}
        style={{ borderRadius: 50 } as ImageStyle}
        className={`rounded-full ${size}`}
      />
    </View>
  );
}

