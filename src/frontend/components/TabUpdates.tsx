import React from "react";
import { View, Text, ScrollView, Dimensions } from "react-native";
import CardAttFollowedProject from "./CardAttFollowedProjects";
import CardAttSinergy from "./CardAttSinergy";
import moment from 'moment';

const window = Dimensions.get('window');

export const TabUpdates = ({notifications}) => {
  
    return (
      <ScrollView style={{ flex: 1, backgroundColor: '#FFF' }}>
        <View style={{ alignItems: 'center', justifyContent: 'center'}}>
          {Object.keys(notifications).map(date => {
            let title;
            if (moment().format('YYYY-MM-DD') === date) {
              title = 'Hoje';
            } else if (moment().subtract(1, 'days').format('YYYY-MM-DD') === date) {
              title = 'Ontem';
            } else {
              title = `Dia ${moment(date).format('DD/MM/YYYY')}`;
            }
  
  
            return (
              <View key={date}>
                <Text style={{fontWeight: 500, fontSize: 18, padding: window.height * 0.01}}>{title}</Text>
                {notifications[date].map((update, index) => (
                  update.type === "Atualização" ? 
                  <CardAttSinergy key={index} {...update} /> :
                  <CardAttFollowedProject key={index} {...update} />
                ))}
              </View>
            );
          })}
        </View>
      </ScrollView>
    );
  };
  