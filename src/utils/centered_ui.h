#pragma once
#include<QScreen>

template<typename T>
void centered_ui(T *windows) {
    QRect screen = QGuiApplication::primaryScreen()->availableGeometry();
    auto size = windows -> geometry();
    windows -> move((screen.width() - size.width()) / 2, (screen.height() - size.height()) / 2);
}