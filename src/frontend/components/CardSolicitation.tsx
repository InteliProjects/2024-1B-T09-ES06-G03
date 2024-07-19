import React, { useState } from 'react';
import { Text, View, Image, Dimensions, TouchableOpacity, TouchableHighlight } from 'react-native';
import Category from '../components/Category';
import { LinearGradient } from 'expo-linear-gradient';
import moment from 'moment';
import WhiteIcon from '../assets/WhiteNotificationIcon.svg'
import { Avatar } from './Avatar';
import CeoAvatar from './CeoAvatar';
import { ceoApi, projectApi } from '../services/api';
import axios from "axios";


const window = Dimensions.get('window');

function updateSynergy(related_synergy_id, id, answer) {

    const getSynergy = async (synergy_id) => {  
        try {
            const response = await projectApi.get(`/synergies/${synergy_id}`);
            return response.data;
    } catch (err) {
        console.error('Error fetching synergy', err);
    }     
    }
    
    const putSynergy = async (synergyData, synergy_id) => {
      try{
        const response = await projectApi.put(`/synergies/${synergy_id}`,
            synergyData   
           );
           return response.data;
      } catch (err) {
            console.error('Error updating synergy', err);
      }
       
    }

    const deleteNotification = async (notification_id) => {
        try {
            const response = await ceoApi.delete(`/notifications/${notification_id}`);
            return response.data;
        } catch (err) {
            console.error('Error deleting notification', err);
        }
        
    }

    const fetchSynergy = async () => {
        const data = await getSynergy(related_synergy_id);
        const updatedSynergy = {};
        updatedSynergy['source_project_id'] = data.source_project.id;
        updatedSynergy['target_project_id'] = data.target_project.id;
        answer === "aceito"?  updatedSynergy['status'] = 'Em andamento': updatedSynergy['status'] = 'Rejeitado';
        updatedSynergy['type'] = data.type;     
        const response = await putSynergy(updatedSynergy, related_synergy_id);
        console.log("resposta do fecthSynergy", response) 
        const response2 = await deleteNotification(id);
        console.log("resposta do deleteNotification", response2)
    }
      
    fetchSynergy();   
}

const CardSolicitation = ({ project, title,  message, data, related_synergy_id, id }) => {


    const onPressAceitar = () => {
      
        updateSynergy(related_synergy_id, id, "aceito");
    }
    
    const onPressRecusar = () => {
        updateSynergy(related_synergy_id, id, "rejeitado");
    }
    // const formattedDate = formatDate(data);
    return (
        <View>
            <View style={{
                padding: 10,
                backgroundColor: '#3A8A88',
                borderTopEndRadius: 10,
                borderTopStartRadius: 10,
                width: window.width * 0.9,
                marginTop: window.height * 0.01,
            }}>
                <View style={{ flexDirection: 'column' }}>
                    <View style={{ flexDirection: 'row', alignItems: 'center' }}>
                        <WhiteIcon width={10} height={10}/>
                        <Text style={{
                            fontWeight: 'bold',
                            fontSize: 10,
                            marginLeft: 5,
                            color: '#fff'
                        }}>
                            {/* {interest} */}
                            Interresse em Integrar BOVAER
                        </Text>
                    </View>
                    <View style={{ flexDirection: 'row', alignItems: 'center', marginTop: 5 }}>
                       <CeoAvatar size='w-12 h-12' link=''/>
                        <View style={{ marginLeft: 10 }}>
                            <Text style={{ fontSize: 18, fontWeight: 500, color: '#fff' }}>
                                {project.name}           
                            </Text>
                            <Text style={{ fontSize: 14, fontWeight: 600, color: '#A7CBCA' }}>{title}</Text>
                            <Text style={{ fontSize: 14, color: '#A7CBCA' }}>
                                {/* {theme} */}
                                Carbonística
                            </Text>
                            <Text style={{ fontSize: 12, color: '#A7CBCA' }}>
                                {/* {email} */}
                                example@email.com
                            </Text>
                        </View>
                        <View style={{
                            marginStart: window.width * 0.05,
                            flexDirection: 'column',
                            alignItems: 'center'
                        }}>
                            <View className={'flex items-center justify-center rounded-full w-[33.5px] h-[33.5px] border-[0.5px] border-white'}>
                                <Category category={"conservation"} circleSize={'w-[32px] h-[32px]'} iconSize={20} />
                            </View>
                            <TouchableOpacity>
                                <Text style={{
                                    marginBottom: 5,
                                    color: '#fff',
                                    fontWeight: 'bold',
                                    fontSize: 10,
                                    marginTop: 10,
                                    backgroundColor: '#679897',
                                    borderRadius: 20,
                                    paddingHorizontal: 10,
                                    paddingVertical: 5
                                }}>Ver Projeto</Text>
                            </TouchableOpacity>
                        </View>
                    </View>
                </View>
            </View>
            <View style={{
                borderBottomEndRadius: 10,
                borderBottomStartRadius: 10,
                backgroundColor: '#f6f6f6',
                width: window.width * 0.9,
                padding: 10,
                marginBottom: window.height * 0.01,
            }}>
                <Text style={{ padding: window.width * 0.04 }}>{message}</Text>
                <View style={{
                    marginHorizontal: window.width * 0.05,
                    flexDirection: 'row',
                    marginTop: window.width * 0.02,
                    alignItems: 'center',
                    justifyContent: 'space-between'
                }}>
                    <TouchableOpacity style={{
                        overflow: 'hidden',
                        height: window.height * 0.05,
                    }}
                    onPress={onPressRecusar}
                    >
                        <View style={{
                            borderRadius: 20,
                            width: window.width * 0.3,
                            height: window.height * 0.04,
                            alignItems: 'center',
                            justifyContent: 'center',
                            backgroundColor: 'transparent',
                            borderColor: '#3A8A88',
                            borderWidth: 2,
                        }}>
                            <Text style={{ color: '#3A8A88', fontSize: 16, fontWeight: 'bold' }}>Recusar</Text>
                        </View>
                    </TouchableOpacity>
                    <TouchableOpacity style={{
                        overflow: 'hidden',
                        height: window.height * 0.05,
                    }}
                    onPress={onPressAceitar}
                    >
                        <LinearGradient
                            colors={['#14514F', '#3A8A88']}
                            start={{ x: 0, y: 0 }}  end={{ x: 1, y: 0 }}
                            style={{
                                borderRadius: 20,
                                width: window.width * 0.4,
                                height: window.height * 0.04,
                                alignItems: 'center',
                                justifyContent: 'center',
                            }}
                        >
                            <Text style={{ color: '#fff', fontSize: 16, fontWeight: 'bold' }}>Aceitar</Text>
                        </LinearGradient>
                    </TouchableOpacity>
                </View>
            </View>
        </View>
    );
};


const formatDate = (dateString) => {
    const date = moment(dateString, 'DD/MM/YYYY');
    const today = moment();
    const yesterday = moment().subtract(1, 'days');
  
    if(date.isSame(today, 'day')){
      return 'Hoje';
    } else if(date.isSame(yesterday, 'day')){
      return 'Ontem';
    } else {
      return date.format('DD/MM/YYYY');
    }
  };



export default CardSolicitation;
