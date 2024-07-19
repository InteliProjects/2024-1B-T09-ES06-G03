import React from "react";
import { View, Text, ScrollView, Dimensions } from "react-native";
import { TabView, SceneMap, TabBar } from 'react-native-tab-view';
import CardAttFollowedProject from "../components/CardAttFollowedProjects";
import CardAttSinergy from "../components/CardAttSinergy";
import CardSolicitation from "../components/CardSolicitation";
import moment from 'moment';
import { coreApi, projectApi, ceoApi } from "../services/api";
import { useEffect, useState } from "react";
import { serializer } from "../metro.config";
import { TabUpdates } from "../components/TabUpdates";
import { TabSolicitations } from "../components/TabSolicitations";



function projectNofications(notifications, projects) {
  const notificationsProjects = notifications.map(notification => {
    const project = projects.filter(project => project.id === notification.related_project_id);
    return { ...notification, project };
  }
  );
  return notificationsProjects;
}

const formatDate = (isoString) => {
  return moment(isoString).format('YYYY-MM-DD');
};

const groupByDate = (cards) => {
  const groups = {};
  cards.forEach(card => {
    const date = formatDate(card.created_at);
    if (!groups[date]) {
      groups[date] = [];
    }
    groups[date].push(card);
  });
  return groups;
};


// const updates = [
//   {
//     interest: "Interesse em Integrar com BOVAER",
//     project: "Way Carbon",
//     name: "Guilherme Vasconcellos",
//     theme: "Carbonística",
//     avatar: "link_para_imagem",
//     status: "Recusado",
//     category: "conservation",
//     data: "25/05/2024",
//     type: "CardAttSinergy"
//   },
//   {
//     title: "Novidades",
//     news: "20 novos projetos entraram no Sr. Match!",
//     data: "25/05/2024",
//     type: "CardAttFollowedProject"
//   },
//   {
//     interest: "Interesse em Integrar com BOVAER",
//     project: "Way Carbon",
//     name: "Guilherme Vasconcellos",
//     theme: "Carbonística",
//     avatar: "link_para_imagem",
//     status: "Aceito",
//     category: "conservation",
//     data: "24/05/2024",
//     type: "CardAttSinergy"
//   },
//   {
//     title: "Novidades",
//     news: "20 novos projetos entraram no Sr. Match!",
//     data: "23/05/2024",
//     type: "CardAttFollowedProject"
//   },
//   // ... outros updates
// ];



export default function Notification({ navigation }) {
  const [index, setIndex] = React.useState(0);
  const [routes] = React.useState([
    { key: 'updates', title: 'Atualizações' },
    { key: 'solicitations', title: 'Solicitações' },
  ]);
  const [notifications, setNotifications] = useState([]);
  const [error, setError] = useState(null);
  const [userId, setUserId] = useState(2);

  useEffect(() => {
    const fetchUserProjects = async () => {
      try{
        const response = await projectApi(`/projects/me`);
        setUserId(response.data[0].user_id)
      } catch (error) {
        console.error('Error fetching projetos:', error);
      }
    }

    const fetchNotifications = async () => {
      try {
        const response = await ceoApi.get(`/notifications/user/${userId}`);
        setNotifications(response.data);
      } catch (error) {
        setError(error);
      }
    };

    fetchUserProjects();
    fetchNotifications();
  }, []);


  console.log("erro", error)
  console.log("notifications in screen", notifications);
  const groupedCards = groupByDate(notifications);

  const renderUpdates = () => {
    return <TabUpdates notifications={groupedCards} />
  }

  const renderSolicitations = () => {
    return <TabSolicitations notifications={groupedCards} />
  }
  const renderScene = SceneMap({
    updates: renderUpdates,
    solicitations: renderSolicitations,
  });

  const renderTabBar = props => (
    <TabBar
      {...props}
      indicatorStyle={{ backgroundColor: '#3A8A88', width: '30%', position: 'absolute', bottom: 0, left: '10%' }}
      style={{ backgroundColor: '#fff', alignContent: 'center', justifyContent: 'center' }}
      labelStyle={{ color: '#000000', fontSize: 14, fontWeight: 'bold' }}

    />
  );

  return (
    <TabView
      navigationState={{ index, routes }}
      renderScene={renderScene}
      renderTabBar={renderTabBar}
      onIndexChange={setIndex}
    />
  );
}