import React from 'react';
import { Text, View, Image, Dimensions } from 'react-native';
import Category from '../components/Category';
import moment from 'moment';
import CeoAvatar from './CeoAvatar';
import BlackNotificationIcon from '../assets/blackNotificationIcon.svg';

const window = Dimensions.get('window');

const CardAttSinergy = ({ interest, project, category_name, subcategory_name, avatar, status, category, data }) => {
    // Determine color based on status
    const color = status === 'Aceito' ? '#3A8A88' : '#BB3756';
  
    // const formattedDate = formatDate(data);

    return (
        <View style={{
            flexDirection: 'row',
            justifyContent: 'space-between',
            alignItems: 'center',
            paddingHorizontal: 20,
            paddingVertical: 10,
            backgroundColor: '#F6F6F6',
            borderRadius: 20,
            width: window.width * 0.9,
            // Add border on one side only
            borderRightWidth: 4,
            borderRightColor: color,
            // Add rounded border on the right side
            borderTopRightRadius: 10,
            borderBottomRightRadius: 10,
            marginTop: window.height * 0.01,
            marginBottom: window.height * 0.01,
            shadowColor: 'rgba(0, 0, 0, 0.60)',
            shadowOpacity: 0.25,
            shadowOffset: { width: 0, height: 4 },
            elevation: 4,
            shadowRadius: 8
        }}>
            <View style={{ flexDirection: 'column' }}>
                <View style={{ flexDirection: 'row', alignItems: 'center', justifyContent: 'space-between', width: '100%' }}>
                    <View className='w-[70%]' style={{ flexDirection: 'row', alignItems: 'center' }}>
                        <BlackNotificationIcon width={10} height={10} />
                        <Text style={{ fontWeight: 500, fontSize: 12, marginLeft: 5 }}>
                            Interresse em integrar com BOVAER
                        </Text>
                    </View>
                    <Text style={{ marginBottom: 5, color: color, fontWeight: 'bold', fontSize: 14, alignSelf: 'flex-end' }}>{status}</Text>
                </View>
                <View style={{ flexDirection: 'row', alignItems: 'center', marginTop: 5, justifyContent: 'space-between' }}>
                    <View style={{ flexDirection: 'row', alignItems: 'center' }}>
                        <CeoAvatar size='w-12 h-12' link='' />
                        <View style={{ marginLeft: 10 }}>
                            <Text style={{ fontSize: 18, fontWeight: 500 }}>{project.name}</Text>
                            <Text style={{ fontSize: 14, fontWeight: 600, color: "#787878" }}>Guilherme Vasconselos</Text>
                            <Text style={{ fontSize: 14 }}>{subcategory_name}</Text>
                        </View>
                    </View>
                    <View style={{ alignItems: 'center' }}>
                        <Category category={category_name} circleSize={'w-[32px] h-[32px]'} iconSize={20} />
                    </View>
                </View>
            </View>
        </View>
    );
}

const formatDate = (dateString) => {
    const date = moment(dateString, 'DD/MM/YYYY');
    const today = moment();
    const yesterday = moment().subtract(1, 'days');

    if (date.isSame(today, 'day')) {
        return 'Hoje';
    } else if (date.isSame(yesterday, 'day')) {
        return 'Ontem';
    } else {
        return date.format('DD/MM/YYYY');
    }
};

export default CardAttSinergy;