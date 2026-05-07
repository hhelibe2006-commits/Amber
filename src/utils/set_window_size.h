#pragma once
#include<QScreen>

template<typename T>
void set_window_size(T *windows, double ratio) {
    QRect rect = QGuiApplication::primaryScreen()->geometry();
    int width = rect.width();
    int height = rect.height();
    windows -> resize(width * ratio, height * ratio);
}